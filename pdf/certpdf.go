package pdf

import (
	"fmt"
	"gencert/cert"
	"github.com/jung-kurt/gofpdf"
	"os"
	path2 "path"
)

type PdfSaver struct {
	OutputDir string
}

func New(outpudir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outpudir, os.ModePerm)
	if err != nil {
		return p, err
	}
	p = &PdfSaver{
		OutputDir: outpudir,
	}
	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")

	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// save file
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path2.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to '%v'\n", path)
	return nil
}
