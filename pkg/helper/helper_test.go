package helper

import (
	"fmt"
	"testing"
)

func TestHelper(t *testing.T) {
	token, _ := GenerateToken("1", "admin")
	//userClaim, _ := AnalyseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6IiIsIm5hbWUiOiJhZG1pbiJ9.z_oqGnWGhVgrp8746TIc2kmHSMhYU0HZ3Rk1eSU_mxU")
	fmt.Println(token)
	//fmt.Println(userClaim, userClaim.Name)
}

func TestMD5(t *testing.T) {
	fmt.Println(GetMd5("123456"))
}
