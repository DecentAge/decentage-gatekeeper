package testsuite

import "errors"

const (
	FakeAdminRole          = "role:admin"
	FakeAdminURL           = "/admin"
	FakeAdminAllURL        = FakeAdminURL + "*"
	FakeAuthAllURL         = "/auth_all/*"
	FakeClientID           = "test"
	FakeSecret             = "test"
	FakeTestAdminRolesURL  = "/test_admin_roles"
	FakeTestRole           = "role:test"
	FakeTestURL            = "/test"
	FakeTestRoleURL        = "/test_role"
	FakeTestWhitelistedURL = "/auth_all/white_listed*"
	TestProxyAccepted      = "Proxy-Accepted"
	ValidUsername          = "test"
	ValidPassword          = "test"
	testEncryptionKey      = "ZSeCYDUxIlhDrmPpa1Ldc7il384esSF2"
	randomLocalHost        = "127.0.0.1:0"
	FakeSignature          = ".SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	FakeCertFilePrefix     = "/gateadmin_crt_"
	FakePrivFilePrefix     = "/gateadmin_priv_"
	FakeCaFilePrefix       = "/gateadmin_ca_"
)

var ErrCreateFakeProxy = errors.New("failed to create fake proxy service")
