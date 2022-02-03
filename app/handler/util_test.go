package handler

/* func TestNewSha256(t *testing.T) {
	for i, tt := range []struct {
		in  []byte
		out string
	}{
		{[]byte("hola"), "b221d9dbb083a7f33428d7c2a3c3198ae925614d70210e28716ccaa7cd4ddb79"},
		{[]byte("adios"), "d8542114d7d40f3c82fc0919efc644df30f4e827c2bd6b83b9dbec8358b2fbc4"},
		{[]byte("password"), "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := newSha256(tt.in)
			if hex.EncodeToString(result) != tt.out {
				t.Errorf("want %v; got %v", tt.out, hex.EncodeToString(result))
			}
		})
	}

} */
