package service

/* func TestMakeGetAllUsersEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  getAllUsersRequest
		out string
	}{
		{getAllUsersRequest{}, ""},
		{getAllUsersRequest{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string
			svc := GetServiceDB()

			//OpenDB
			err := svc.OpenDB(DBDriver, PsqlInfo)
			if err != nil {
				t.Error(err)
			}
			defer svc.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := svc.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			r, err := MakeGetAllUsersEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(getAllUsersResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
} */

/* func TestMakeGetUserByIDEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  getUserByIDRequest
		out string
	}{
		{getUserByIDRequest{1}, ""},
		{getUserByIDRequest{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string
			svc := GetServiceDB()

			//OpenDB
			err := svc.OpenDB(DBDriver, PsqlInfo)
			if err != nil {
				t.Error(err)
			}
			defer svc.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := svc.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			r, err := MakeGetUserByIDEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(getUserByIDResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
} */

/* func TestMakeGetUserByUsernameAndPasswordEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  getUserByUsernameAndPasswordRequest
		out string
	}{
		{getUserByUsernameAndPasswordRequest{"cesar", "01234"}, ""},
		{getUserByUsernameAndPasswordRequest{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string
			svc := GetServiceDB()

			//OpenDB
			err := svc.OpenDB(DBDriver, PsqlInfo)
			if err != nil {
				t.Error(err)
			}
			defer svc.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := svc.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			r, err := MakeGetUserByUsernameAndPasswordEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(getUserByUsernameAndPasswordResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
} */

/* func TestGetIDByUsernameEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  getIDByUsernameRequest
		out string
	}{
		{getIDByUsernameRequest{"cesar"}, ""},
		{getIDByUsernameRequest{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string
			svc := GetServiceDB()

			//OpenDB
			err := svc.OpenDB(DBDriver, PsqlInfo)
			if err != nil {
				t.Error(err)
			}
			defer svc.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := svc.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			r, err := MakeGetIDByUsernameEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(getIDByUsernameResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
} */

/* func TestMakeInsertUserEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  insertUserRequest
		out string
	}{
		{insertUserRequest{}, ""},
		{insertUserRequest{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string
			svc := GetServiceDB()

			//OpenDB
			err := svc.OpenDB(DBDriver, PsqlInfo)
			if err != nil {
				t.Error(err)
			}
			defer svc.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := svc.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			r, err := MakeInsertUserEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(insertUserResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
} */

/* func TestMakeDeleteUserEndpoint(t *testing.T) {
	for i, tt := range []struct {
		in  deleteUserRequest
		out string
	}{
		{deleteUserRequest{}, ""},
		{deleteUserRequest{}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// var resultErr string
			svc := GetServiceDB()

			//OpenDB
			err := svc.OpenDB(DBDriver, PsqlInfo)
			if err != nil {
				t.Error(err)
			}
			defer svc.db.Close()

			// generate confict closing db
			if tt.out == "database is closed" {
				err := svc.db.Close()
				if err != nil {
					t.Error(err)
				}
			}

			r, err := MakeDeleteUserEndpoint(svc)(context.TODO(), tt.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(deleteUserResponse)
			if !ok {
				t.Error("response is not of the type indicated")
			}

			if !strings.Contains(result.Err, tt.out) {
				t.Errorf("want %v; got %v", tt.out, result.Err)
			}
		})
	}
} */
