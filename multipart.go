package idea

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"os"

	geodata "jpmenezes.com/idebo/gen/geodata"
)

// GeodataUploadDecoderFunc implements the multipart decoder for service
// "geodata" endpoint "upload". The decoder must populate the argument p after
// encoding.
func GeodataUploadDecoderFunc(mr *multipart.Reader, p **geodata.FilesUpload) error {
	res := geodata.FilesUpload{}

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}

		_, params, err := mime.ParseMediaType(p.Header.Get("Content-Disposition"))
		if err != nil {
			// can't process this entry, it probably isn't an image
			continue
		}

		disposition, _, err := mime.ParseMediaType(p.Header.Get("Content-Type"))
		if err != nil {
			log.Println(err)
		}
		// the disposition can be, for example 'image/jpeg' or 'video/mp4'
		// I want to support only image files!
		// if err != nil || !strings.HasPrefix(disposition, "image/") {
		// 	// can't process this entry, it probably isn't an image
		// 	continue
		// }

		if params["name"] == "files" {
			bytes, err := ioutil.ReadAll(p)
			if err != nil {
				// can't process this entry, for some reason
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			filename := params["filename"]
			imageUpload := geodata.FileUpload{
				Type:  &disposition,
				Bytes: bytes,
				Name:  &filename,
			}
			res.Files = append(res.Files, &imageUpload)
		}
	}
	*p = &res
	return nil
}

// GeodataUploadEncoderFunc implements the multipart encoder for service
// "geodata" endpoint "upload".
func GeodataUploadEncoderFunc(mw *multipart.Writer, p *geodata.FilesUpload) error {
	// Add multipart request encoder logic here
	return nil
}
