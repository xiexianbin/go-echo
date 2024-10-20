package app

// func TestDecodeQuery(t *testing.T) {
// 	type args struct {
// 		rawQuery string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *SearchOptions
// 		wantErr bool
// 	}{
// 		{
// 			name: "order-sort",
// 			args: args{
// 				rawQuery: "q=search+term&sort=name,asc&sort=created_at,desc",
// 			},
// 			want:    &SearchOptions{},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := DecodeQuery(tt.args.rawQuery)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("DecodeQuery() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("DecodeQuery() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
