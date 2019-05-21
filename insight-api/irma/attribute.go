// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package irma

type Attribute string

// TODO: Create a more structured mapping for the attribute value, perhaps based on the following.
// TODO: See if IRMA doesn't already have an attribute implementation which can be used.
// TODO: Add JSON marshal/unmarshal functions for these
// type Attribute struct {
// 	Credential
// 	AttributeName string
// }

// func (a Attribute) String() string {
// 	return a.Credential.String() + `.` + a.AttributeName
// }

// type Credential struct {
// 	Issuer
// 	CredentialName string
// }

// func (c Credential) String() string {
// 	return c.Issuer.String() + `.` + c.CredentialName
// }

// type Issuer struct {
// 	SchemeManager
// 	IssuerName string
// }

// func (i Issuer) String() string {
// 	return i.SchemeManager.String() + `.` + i.IssuerName
// }

// type SchemeManager struct {
// 	SchemeManagerName string
// }

// func (s SchemeManager) String() string {
// 	return s.SchemeManagerName
// }
