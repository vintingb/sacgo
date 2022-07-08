package sac

import (
	"github.com/vintingb/sacgo/line"
	"gonum.org/v1/plot/vg"
	"net/http"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	fp, _ := os.Open("test.SAC")
	s := Decode(fp)
	t.Log(s.SacHead)
	t.Log(s.SacData.Pts[0])
}

func TestHas(t *testing.T) {
	fp, _ := os.Open("test.SAC")
	s := Decode(fp)
	t.Log(s.SacHead.Has("delta"))
	t.Log(s.SacHead.Has("delt"))
}

func TestPlot(t *testing.T) {
	fp, _ := os.Open("test.SAC")
	s := Decode(fp)
	line.Delta = float64(s.SacHead.Delta)
	p := s.SacData.plot("test")
	p.Save(20*vg.Inch, 5*vg.Inch, "test.pdf")
}

func TestWeb(t *testing.T) {
	http.HandleFunc("/sac", func(w http.ResponseWriter, r *http.Request) {
		fp, _ := os.Open("test.SAC")
		s := Decode(fp)
		line.Delta = float64(s.SacHead.Delta)
		p, _ := s.Plot("Test")
		p.WriteTo(w)
	})
	http.ListenAndServe(":3000", nil)
}
