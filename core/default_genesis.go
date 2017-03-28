// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
    "compress/gzip"
    "encoding/base64"
    "io"
    "strings"
)

func NewDefaultGenesisReader() (io.Reader, error) {
    return gzip.NewReader(base64.NewDecoder(base64.StdEncoding, strings.NewReader(defaultGenesisBlock)))
}

// defaultGenesisBlock is a gzip compressed dump of the official default Dubaicoin
// genesis block.

const ( defaultGenesisBlock = "H4sIAAAAAAAA/62RTU7DMBCF180pIq+7sGMndtiBIsSCS4ztGWopP1VjpKAqd8ckDULQSggxO78375uxfc7yVKwfeofsLmd84l9LCs32a0sMHY4RuuPaVprGNI/lw2Yf4YR9fILxcAXzh9q4OMUTNBDhgt30FxifQxfiKgtT3X9aPhAF99rGt9U035BdmA7/vqcbQm9hvPaIv0hD2w4uRc/LcZHSpbwilK4mJI5GEnlhUZAkR0aVTsuiLgvA+iOXMwstXD6xEjemzftst0tkzxNNKesMWF0Al+BVZWQtVFE7gail0oDW/yDfBC97z9n8DtDgxIlQAgAA"

defaultTestnetGenesisBlock = defaultGenesisBlock

)