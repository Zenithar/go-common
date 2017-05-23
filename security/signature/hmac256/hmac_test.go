package hmac256_test

import (
	"testing"

	"go.zenithar.org/common/security/signature/hmac256"
)

func TestSign(t *testing.T) {
	secret := "toto"
	payload := `{"event":"agent:result","payload":{"ssdeep":"3072:Mi3sSMsKeDYlsSoDZIa69fu2hwQ2yLi0wVraJW/hgq5BxxbkGXwxxDlamBQky:j3sSTKoYlfoDoukwQ7LEhzHjbjAzpBdy","trid":["88.7% (.DOCX) Word Microsoft Office Open XML Format document (31500/1/5)","11.2% (.ZIP) ZIP compressed archive (4000/1)"],"exiftool":{"ExifTool Version Number":"10.15","File Size":"189 kB","File Type":"ZIP","File Type Extension":"zip","MIME Type":"application/zip","Zip Bit Flag":"0x0006","Zip CRC":"0xbc96d396","Zip Compressed Size":"377","Zip Compression":"Deflated","Zip File Name":"[Content_Types].xml","Zip Required Version":"20","Zip Uncompressed Size":"1546"}},"agent":"fileinfo","analyse_id":"12345678","timestamp":1465983500}`
	nonce := "1465983500"
	analyse_id := "12345678"

	sig := hmac256.GetSignature([]byte(secret), []byte(payload), []byte("fileinfo"), []byte(analyse_id), []byte(nonce))
	if sig != "iGNzl386q9X0kEzI3QeI4943ZBIiDuknGA8vybQSuIY=" {
		t.Errorf("Signature is not as expected ! %s", sig)
		t.FailNow()
	}
}

func TestCompare(t *testing.T) {
	valid, err := hmac256.CompareSignatures("1XWot6XtD9_FRuW4brK0Dl76Gnu68HdCqGayMdsbjco=", "1XWot6XtD9_FRuW4brK0Dl76Gnu68HdCqGayMdsbjco=")
	if err != nil {
		t.Error("It should not have error")
		t.FailNow()
	}
	if !valid {
		t.Error("Comparison should be good !")
		t.FailNow()
	}

	valid, err = hmac256.CompareSignatures("", "1XWot6XtD9_FRuW4brK0Dl76Gnu68HdCqGayMdsbjco=")
	if err != nil {
		t.Error("It should not have error")
		t.FailNow()
	}
	if valid {
		t.Error("Comparison should be invalid !")
		t.FailNow()
	}
}
