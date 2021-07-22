package main

import "testing"

func TestCircle_Area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name:    "return error if negative or null value",
			fields:  fields{radius: 0},
			want:    0,
			wantErr: true,
		},
		{
			name:    "return area of circle",
			fields:  fields{radius: 8},
			want:    201.06192982974676,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Circle{
				radius: tt.fields.radius,
			}
			got, err := s.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Area() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimetr(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name:    "return error if negative or null value",
			fields:  fields{radius: 0},
			want:    0,
			wantErr: true,
		},
		{
			name:    "return perimetr of circle",
			fields:  fields{radius: 8},
			want:    50.26548245743669,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Circle{
				radius: tt.fields.radius,
			}
			got, err := s.Perimetr()
			if (err != nil) != tt.wantErr {
				t.Errorf("Perimetr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Perimetr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name:    "error if negative or null values",
			fields:  fields{height: -1, width: -1},
			want:    0,
			wantErr: true,
		},
		{
			name:    "area of rectangle",
			fields:  fields{height: 9, width: 3},
			want:    27,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			got, err := s.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Area() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimetr(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name:    "error if negative or null values",
			fields:  fields{height: -1, width: -1},
			want:    0,
			wantErr: true,
		},
		{
			name:    "perimetr of rectangle",
			fields:  fields{height: 9, width: 3},
			want:    24,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			got, err := s.Perimetr()
			if (err != nil) != tt.wantErr {
				t.Errorf("Perimetr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Perimetr() got = %v, want %v", got, tt.want)
			}
		})
	}
}
