package ocpp_test

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/road-labs/ocpp-types-go/gen/ocpp15"
)

// TestAuthorizeRoundTrip proves the same struct value serialises to both an OCPP-J JSON payload and an OCPP-S SOAP
// element for the CentralSystemService namespace, and unmarshals back from either.
func TestAuthorizeRoundTrip(t *testing.T) {
	req := ocpp15.Authorize{IDTag: "ABC123"}

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("json marshal: %v", err)
	}
	if got, want := string(jsonPayload), `{"idTag":"ABC123"}`; got != want {
		t.Errorf("json payload: got %q, want %q", got, want)
	}

	xmlPayload, err := xml.Marshal(req)
	if err != nil {
		t.Fatalf("xml marshal: %v", err)
	}
	wantXML := `<authorizeRequest xmlns="urn://Ocpp/Cs/2012/06/"><idTag>ABC123</idTag></authorizeRequest>`
	if string(xmlPayload) != wantXML {
		t.Errorf("xml payload:\n got  %s\n want %s", string(xmlPayload), wantXML)
	}

	var backFromJSON ocpp15.Authorize
	if err := json.Unmarshal(jsonPayload, &backFromJSON); err != nil {
		t.Fatalf("json unmarshal: %v", err)
	}
	if backFromJSON.IDTag != req.IDTag {
		t.Errorf("json round trip: IDTag mismatch got %q want %q", backFromJSON.IDTag, req.IDTag)
	}

	var backFromXML ocpp15.Authorize
	if err := xml.Unmarshal(xmlPayload, &backFromXML); err != nil {
		t.Fatalf("xml unmarshal: %v", err)
	}
	if backFromXML.IDTag != req.IDTag {
		t.Errorf("xml round trip: IDTag mismatch got %q want %q", backFromXML.IDTag, req.IDTag)
	}
}

// TestMeterValuesAttributeMode covers the OCPP 1.5 quirk where the sampled-value metadata (context, format, measurand,
// location, unit) are XML *attributes* on the value element and the reading itself is character data — while JSON
// encodes all six as sibling object properties.
func TestMeterValuesAttributeMode(t *testing.T) {
	ctx := ocpp15.ContextEnumTypeSamplePeriodic
	measurand := ocpp15.MeasurandEnumTypeEnergyActiveImportRegister
	unit := ocpp15.UnitEnumTypeWh

	mv := ocpp15.MeterValues{
		ConnectorID: 1,
		Values: []ocpp15.MeterValuesValuesElem{{
			Timestamp: time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC),
			Value: []ocpp15.MeterValueType{{
				Value:     "12345",
				Context:   &ctx,
				Measurand: &measurand,
				Unit:      &unit,
			}},
		}},
	}

	xmlPayload, err := xml.Marshal(mv)
	if err != nil {
		t.Fatalf("xml marshal: %v", err)
	}
	// The value element must carry attributes for the metadata and character data for the reading.
	xmlText := string(xmlPayload)
	for _, want := range []string{
		`xmlns="urn://Ocpp/Cs/2012/06/"`,
		`context="Sample.Periodic"`,
		`measurand="Energy.Active.Import.Register"`,
		`unit="Wh"`,
		`>12345<`,
	} {
		if !strings.Contains(xmlText, want) {
			t.Errorf("xml payload missing %q; got:\n%s", want, xmlText)
		}
	}

	jsonPayload, err := json.Marshal(mv)
	if err != nil {
		t.Fatalf("json marshal: %v", err)
	}
	// JSON encodes all metadata as siblings of value.
	jsonText := string(jsonPayload)
	for _, want := range []string{
		`"value":"12345"`,
		`"context":"Sample.Periodic"`,
		`"measurand":"Energy.Active.Import.Register"`,
		`"unit":"Wh"`,
	} {
		if !strings.Contains(jsonText, want) {
			t.Errorf("json payload missing %q; got:\n%s", want, jsonText)
		}
	}

	// Round-trip XML back and confirm the metadata survived the attribute encoding.
	var back ocpp15.MeterValues
	if err := xml.Unmarshal(xmlPayload, &back); err != nil {
		t.Fatalf("xml unmarshal: %v", err)
	}
	if len(back.Values) != 1 || len(back.Values[0].Value) != 1 {
		t.Fatalf("unexpected structure: %#v", back)
	}
	sv := back.Values[0].Value[0]
	if sv.Value != "12345" {
		t.Errorf("xml round trip: reading %q, want %q", sv.Value, "12345")
	}
	if sv.Context == nil || *sv.Context != ocpp15.ContextEnumTypeSamplePeriodic {
		t.Errorf("xml round trip: context %v, want %v", sv.Context, ocpp15.ContextEnumTypeSamplePeriodic)
	}
	if sv.Unit == nil || *sv.Unit != ocpp15.UnitEnumTypeWh {
		t.Errorf("xml round trip: unit %v, want %v", sv.Unit, ocpp15.UnitEnumTypeWh)
	}
}

// TestDataTransferDirectionHelpers proves that the bi-directional DataTransfer message picks the correct SOAP namespace
// via the generated FromChargePoint / FromCentralSystem helpers, while still producing an identical JSON payload
// regardless of direction.
func TestDataTransferDirectionHelpers(t *testing.T) {
	messageID := "someMessageId"
	build := func() *ocpp15.DataTransfer {
		return &ocpp15.DataTransfer{VendorID: "com.example.vendor", MessageID: &messageID}
	}

	// The XMLName tag is namespace-less by default, so a raw Marshal without a direction helper
	// leaves the element in no namespace. This is a deliberate signal to callers.
	rawXML, err := xml.Marshal(build())
	if err != nil {
		t.Fatalf("raw xml marshal: %v", err)
	}
	if strings.Contains(string(rawXML), "xmlns=") {
		t.Errorf("raw dataTransfer should have no xmlns; got %s", rawXML)
	}

	cpXML, err := xml.Marshal(build().FromChargePoint())
	if err != nil {
		t.Fatalf("FromChargePoint marshal: %v", err)
	}
	if !strings.Contains(string(cpXML), `xmlns="urn://Ocpp/Cs/2012/06/"`) {
		t.Errorf("FromChargePoint should target CS namespace; got %s", cpXML)
	}

	csXML, err := xml.Marshal(build().FromCentralSystem())
	if err != nil {
		t.Fatalf("FromCentralSystem marshal: %v", err)
	}
	if !strings.Contains(string(csXML), `xmlns="urn://Ocpp/Cp/2012/06/"`) {
		t.Errorf("FromCentralSystem should target CP namespace; got %s", csXML)
	}

	// Direction never leaks into JSON.
	cpJSON, err := json.Marshal(build().FromChargePoint())
	if err != nil {
		t.Fatalf("json marshal after FromChargePoint: %v", err)
	}
	csJSON, err := json.Marshal(build().FromCentralSystem())
	if err != nil {
		t.Fatalf("json marshal after FromCentralSystem: %v", err)
	}
	if string(cpJSON) != string(csJSON) {
		t.Errorf("json payloads should be identical across directions:\n cp: %s\n cs: %s", cpJSON, csJSON)
	}
	if strings.Contains(string(cpJSON), "XMLName") {
		t.Errorf("json payload must not contain XMLName; got %s", cpJSON)
	}
}

// TestBootNotificationWireOrder pins the field order that the XML encoder emits for BootNotification. The WSDL declares
// a specific <s:sequence> that strict SOAP validators enforce; Go's encoding/xml uses struct-declaration order, which
// the post-processor sorts to match the WSDL.
func TestBootNotificationWireOrder(t *testing.T) {
	serial := "SN-123"
	req := ocpp15.BootNotification{
		ChargePointVendor:       "Vendor",
		ChargePointModel:        "Model",
		ChargePointSerialNumber: &serial,
	}
	out, err := xml.Marshal(req)
	if err != nil {
		t.Fatalf("xml marshal: %v", err)
	}
	// Verify strict WSDL sequence order: chargePointVendor, chargePointModel, chargePointSerialNumber, ...
	iVendor := strings.Index(string(out), "<chargePointVendor>")
	iModel := strings.Index(string(out), "<chargePointModel>")
	iSerial := strings.Index(string(out), "<chargePointSerialNumber>")
	if iVendor < 0 || iModel < 0 || iSerial < 0 {
		t.Fatalf("expected all three elements; got:\n%s", out)
	}
	if !(iVendor < iModel && iModel < iSerial) {
		t.Errorf("elements out of WSDL sequence order:\n%s", out)
	}
}

// TestIdTagInfoWireOrder confirms IdTagInfo children arrive in status/expiryDate/parentIdTag order, matching the
// shared complexType sequence in the WSDL.
func TestIdTagInfoWireOrder(t *testing.T) {
	parent := "PARENT"
	expiry := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)
	resp := ocpp15.AuthorizeResponse{
		IDTagInfo: ocpp15.AuthorizeResponseIDTagInfo{
			Status:      ocpp15.AuthorizeResponseIDTagInfoStatusAccepted,
			ExpiryDate:  &expiry,
			ParentIDTag: &parent,
		},
	}
	out, err := xml.Marshal(resp)
	if err != nil {
		t.Fatalf("xml marshal: %v", err)
	}
	iStatus := strings.Index(string(out), "<status>")
	iExpiry := strings.Index(string(out), "<expiryDate>")
	iParent := strings.Index(string(out), "<parentIdTag>")
	if iStatus < 0 || iExpiry < 0 || iParent < 0 {
		t.Fatalf("expected all three IdTagInfo elements; got:\n%s", out)
	}
	if !(iStatus < iExpiry && iExpiry < iParent) {
		t.Errorf("IdTagInfo children out of WSDL sequence order:\n%s", out)
	}
}

// TestSendLocalListWireForm pins the 1.5 SendLocalList wire form on both transports. The 1.5 JSON schema was originally
// published unofficially by an early adopter and diverged from the OCA WSDL in two ways: it dropped the "hash"
// integrity element and Americanised the "localAuthorisationList" element name. Both have since been aligned in-tree so
// a single struct round-trips on both transports without post-processor gymnastics.
func TestSendLocalListWireForm(t *testing.T) {
	hash := "deadbeef"
	req := ocpp15.SendLocalList{
		ListVersion:            1,
		UpdateType:             ocpp15.SendLocalListUpdateTypeFull,
		LocalAuthorisationList: []ocpp15.SendLocalListLocalAuthorisationListElem{{IDTag: "ABC"}},
		Hash:                   &hash,
	}

	xmlOut, err := xml.Marshal(req)
	if err != nil {
		t.Fatalf("xml marshal: %v", err)
	}
	for _, want := range []string{
		"<localAuthorisationList>",
		"<hash>deadbeef</hash>",
	} {
		if !strings.Contains(string(xmlOut), want) {
			t.Errorf("xml payload missing %q; got:\n%s", want, xmlOut)
		}
	}
	if strings.Contains(string(xmlOut), "<localAuthorizationList>") {
		t.Errorf("xml payload must not use Americanised spelling; got:\n%s", xmlOut)
	}

	jsonOut, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("json marshal: %v", err)
	}
	for _, want := range []string{
		`"localAuthorisationList"`,
		`"hash":"deadbeef"`,
	} {
		if !strings.Contains(string(jsonOut), want) {
			t.Errorf("json payload missing %q; got: %s", want, jsonOut)
		}
	}
	if strings.Contains(string(jsonOut), "localAuthorization") {
		t.Errorf("json payload must not use Americanised spelling; got: %s", jsonOut)
	}
}

// TestStopTransactionMeterValueAttributeMode covers a 1.5 quirk where the JSON schema inlines the meter-value shape
// inside StopTransaction.transactionData rather than $ref'ing MeterValueType, so go-jsonschema hoists it into a
// separate anonymous type. The XML wire form must still use the attribute-plus-chardata encoding, matching the WSDL.
func TestStopTransactionMeterValueAttributeMode(t *testing.T) {
	ctx := ocpp15.StopTransactionTransactionDataElemValuesElemValueElemContextSamplePeriodic
	unit := ocpp15.StopTransactionTransactionDataElemValuesElemValueElemUnitWh

	sample := ocpp15.StopTransactionTransactionDataElemValuesElemValueElem{
		Value:   "12345",
		Context: &ctx,
		Unit:    &unit,
	}
	out, err := xml.Marshal(sample)
	if err != nil {
		t.Fatalf("xml marshal: %v", err)
	}
	for _, want := range []string{
		`context="Sample.Periodic"`,
		`unit="Wh"`,
		`>12345<`,
	} {
		if !strings.Contains(string(out), want) {
			t.Errorf("xml payload missing %q; got:\n%s", want, out)
		}
	}
}

// TestChargePointDirectionNamespace confirms that CS-bound and CP-bound messages carry the correct per-direction
// namespace URN.
func TestChargePointDirectionNamespace(t *testing.T) {
	for _, tc := range []struct {
		name string
		msg  any
		ns   string
	}{
		{"Authorize is CS-bound", ocpp15.Authorize{IDTag: "X"}, "urn://Ocpp/Cs/2012/06/"},
		{"Reset is CP-bound", ocpp15.Reset{Type: ocpp15.ResetTypeSoft}, "urn://Ocpp/Cp/2012/06/"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out, err := xml.Marshal(tc.msg)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if !strings.Contains(string(out), `xmlns="`+tc.ns+`"`) {
				t.Errorf("expected xmlns %q in output:\n%s", tc.ns, out)
			}
		})
	}
}
