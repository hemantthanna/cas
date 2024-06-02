package cas

import (
	"context"
	"path/filepath"

	"github.com/hemantthanna/cas/schema"
	"github.com/hemantthanna/cas/types"

	"github.com/hemantthanna/cas/storage"
)

func (s *Storage) indexFileByExt(ctx context.Context, name string) storage.BlobWriter {
	ext := filepath.Ext(name)
	switch ext {
	case ".gz":
		return s.indexGZIP(ctx)
	}
	return nil
}

func (s *Storage) storeIndexByExt(ctx context.Context, name string, orig, ref types.SizedRef) (types.SizedRef, error) {
	ext := filepath.Ext(name)
	switch ext {
	case ".gz":
		m := &schema.Compressed{
			Arch: orig, Ref: ref, Algo: "gzip",
		}
		return s.StoreSchema(ctx, m)
	}
	return types.SizedRef{}, nil
}
