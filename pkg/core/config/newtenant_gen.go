package config

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *AppConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "database_url":
			z.DatabaseURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "smtp":
			err = z.SMTP.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "twilio":
			var zbzg uint32
			zbzg, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zbzg > 0 {
				zbzg--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "account_sid":
					z.Twilio.AccountSID, err = dc.ReadString()
					if err != nil {
						return
					}
				case "auth_token":
					z.Twilio.AuthToken, err = dc.ReadString()
					if err != nil {
						return
					}
				case "from":
					z.Twilio.From, err = dc.ReadString()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "nexmo":
			var zbai uint32
			zbai, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zbai > 0 {
				zbai--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "api_key":
					z.Nexmo.APIKey, err = dc.ReadString()
					if err != nil {
						return
					}
				case "secret":
					z.Nexmo.APISecret, err = dc.ReadString()
					if err != nil {
						return
					}
				case "from":
					z.Nexmo.From, err = dc.ReadString()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *AppConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "database_url"
	err = en.Append(0x84, 0xac, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.DatabaseURL)
	if err != nil {
		return
	}
	// write "smtp"
	err = en.Append(0xa4, 0x73, 0x6d, 0x74, 0x70)
	if err != nil {
		return err
	}
	err = z.SMTP.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "twilio"
	// map header, size 3
	// write "account_sid"
	err = en.Append(0xa6, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x83, 0xab, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Twilio.AccountSID)
	if err != nil {
		return
	}
	// write "auth_token"
	err = en.Append(0xaa, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Twilio.AuthToken)
	if err != nil {
		return
	}
	// write "from"
	err = en.Append(0xa4, 0x66, 0x72, 0x6f, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Twilio.From)
	if err != nil {
		return
	}
	// write "nexmo"
	// map header, size 3
	// write "api_key"
	err = en.Append(0xa5, 0x6e, 0x65, 0x78, 0x6d, 0x6f, 0x83, 0xa7, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Nexmo.APIKey)
	if err != nil {
		return
	}
	// write "secret"
	err = en.Append(0xa6, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Nexmo.APISecret)
	if err != nil {
		return
	}
	// write "from"
	err = en.Append(0xa4, 0x66, 0x72, 0x6f, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Nexmo.From)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AppConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "database_url"
	o = append(o, 0x84, 0xac, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.DatabaseURL)
	// string "smtp"
	o = append(o, 0xa4, 0x73, 0x6d, 0x74, 0x70)
	o, err = z.SMTP.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "twilio"
	// map header, size 3
	// string "account_sid"
	o = append(o, 0xa6, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x83, 0xab, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64)
	o = msgp.AppendString(o, z.Twilio.AccountSID)
	// string "auth_token"
	o = append(o, 0xaa, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e)
	o = msgp.AppendString(o, z.Twilio.AuthToken)
	// string "from"
	o = append(o, 0xa4, 0x66, 0x72, 0x6f, 0x6d)
	o = msgp.AppendString(o, z.Twilio.From)
	// string "nexmo"
	// map header, size 3
	// string "api_key"
	o = append(o, 0xa5, 0x6e, 0x65, 0x78, 0x6d, 0x6f, 0x83, 0xa7, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Nexmo.APIKey)
	// string "secret"
	o = append(o, 0xa6, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	o = msgp.AppendString(o, z.Nexmo.APISecret)
	// string "from"
	o = append(o, 0xa4, 0x66, 0x72, 0x6f, 0x6d)
	o = msgp.AppendString(o, z.Nexmo.From)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AppConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "database_url":
			z.DatabaseURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "smtp":
			bts, err = z.SMTP.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "twilio":
			var zajw uint32
			zajw, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zajw > 0 {
				zajw--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "account_sid":
					z.Twilio.AccountSID, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "auth_token":
					z.Twilio.AuthToken, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "from":
					z.Twilio.From, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "nexmo":
			var zwht uint32
			zwht, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zwht > 0 {
				zwht--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "api_key":
					z.Nexmo.APIKey, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "secret":
					z.Nexmo.APISecret, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "from":
					z.Nexmo.From, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AppConfiguration) Msgsize() (s int) {
	s = 1 + 13 + msgp.StringPrefixSize + len(z.DatabaseURL) + 5 + z.SMTP.Msgsize() + 7 + 1 + 12 + msgp.StringPrefixSize + len(z.Twilio.AccountSID) + 11 + msgp.StringPrefixSize + len(z.Twilio.AuthToken) + 5 + msgp.StringPrefixSize + len(z.Twilio.From) + 6 + 1 + 8 + msgp.StringPrefixSize + len(z.Nexmo.APIKey) + 7 + msgp.StringPrefixSize + len(z.Nexmo.APISecret) + 5 + msgp.StringPrefixSize + len(z.Nexmo.From)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *CORSConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zhct uint32
	zhct, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zhct > 0 {
		zhct--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "origin":
			z.Origin, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z CORSConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "origin"
	err = en.Append(0x81, 0xa6, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Origin)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z CORSConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "origin"
	o = append(o, 0x81, 0xa6, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e)
	o = msgp.AppendString(o, z.Origin)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CORSConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcua uint32
	zcua, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcua > 0 {
		zcua--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "origin":
			z.Origin, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z CORSConfiguration) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.Origin)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *DeploymentRoute) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zdaf uint32
	zdaf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zdaf > 0 {
		zdaf--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "version":
			z.Version, err = dc.ReadString()
			if err != nil {
				return
			}
		case "path":
			z.Path, err = dc.ReadString()
			if err != nil {
				return
			}
		case "type":
			z.Type, err = dc.ReadString()
			if err != nil {
				return
			}
		case "type_config":
			var zpks uint32
			zpks, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.TypeConfig == nil && zpks > 0 {
				z.TypeConfig = make(map[string]interface{}, zpks)
			} else if len(z.TypeConfig) > 0 {
				for key, _ := range z.TypeConfig {
					delete(z.TypeConfig, key)
				}
			}
			for zpks > 0 {
				zpks--
				var zxhx string
				var zlqf interface{}
				zxhx, err = dc.ReadString()
				if err != nil {
					return
				}
				zlqf, err = dc.ReadIntf()
				if err != nil {
					return
				}
				z.TypeConfig[zxhx] = zlqf
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *DeploymentRoute) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "version"
	err = en.Append(0x84, 0xa7, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Version)
	if err != nil {
		return
	}
	// write "path"
	err = en.Append(0xa4, 0x70, 0x61, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Path)
	if err != nil {
		return
	}
	// write "type"
	err = en.Append(0xa4, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Type)
	if err != nil {
		return
	}
	// write "type_config"
	err = en.Append(0xab, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.TypeConfig)))
	if err != nil {
		return
	}
	for zxhx, zlqf := range z.TypeConfig {
		err = en.WriteString(zxhx)
		if err != nil {
			return
		}
		err = en.WriteIntf(zlqf)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DeploymentRoute) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "version"
	o = append(o, 0x84, 0xa7, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Version)
	// string "path"
	o = append(o, 0xa4, 0x70, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.Path)
	// string "type"
	o = append(o, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "type_config"
	o = append(o, 0xab, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67)
	o = msgp.AppendMapHeader(o, uint32(len(z.TypeConfig)))
	for zxhx, zlqf := range z.TypeConfig {
		o = msgp.AppendString(o, zxhx)
		o, err = msgp.AppendIntf(o, zlqf)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DeploymentRoute) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjfb uint32
	zjfb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjfb > 0 {
		zjfb--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "version":
			z.Version, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "path":
			z.Path, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "type_config":
			var zcxo uint32
			zcxo, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.TypeConfig == nil && zcxo > 0 {
				z.TypeConfig = make(map[string]interface{}, zcxo)
			} else if len(z.TypeConfig) > 0 {
				for key, _ := range z.TypeConfig {
					delete(z.TypeConfig, key)
				}
			}
			for zcxo > 0 {
				var zxhx string
				var zlqf interface{}
				zcxo--
				zxhx, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zlqf, bts, err = msgp.ReadIntfBytes(bts)
				if err != nil {
					return
				}
				z.TypeConfig[zxhx] = zlqf
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *DeploymentRoute) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Version) + 5 + msgp.StringPrefixSize + len(z.Path) + 5 + msgp.StringPrefixSize + len(z.Type) + 12 + msgp.MapHeaderSize
	if z.TypeConfig != nil {
		for zxhx, zlqf := range z.TypeConfig {
			_ = zlqf
			s += msgp.StringPrefixSize + len(zxhx) + msgp.GuessSize(zlqf)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *FromScratchOptions) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zeff uint32
	zeff, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zeff > 0 {
		zeff--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "AppName":
			z.AppName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "DatabaseURL":
			z.DatabaseURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "APIKey":
			z.APIKey, err = dc.ReadString()
			if err != nil {
				return
			}
		case "MasterKey":
			z.MasterKey, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *FromScratchOptions) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "AppName"
	err = en.Append(0x84, 0xa7, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.AppName)
	if err != nil {
		return
	}
	// write "DatabaseURL"
	err = en.Append(0xab, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.DatabaseURL)
	if err != nil {
		return
	}
	// write "APIKey"
	err = en.Append(0xa6, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.APIKey)
	if err != nil {
		return
	}
	// write "MasterKey"
	err = en.Append(0xa9, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.MasterKey)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FromScratchOptions) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "AppName"
	o = append(o, 0x84, 0xa7, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.AppName)
	// string "DatabaseURL"
	o = append(o, 0xab, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.DatabaseURL)
	// string "APIKey"
	o = append(o, 0xa6, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.APIKey)
	// string "MasterKey"
	o = append(o, 0xa9, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.MasterKey)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FromScratchOptions) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrsw uint32
	zrsw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrsw > 0 {
		zrsw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "AppName":
			z.AppName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "DatabaseURL":
			z.DatabaseURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "APIKey":
			z.APIKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "MasterKey":
			z.MasterKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FromScratchOptions) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.AppName) + 12 + msgp.StringPrefixSize + len(z.DatabaseURL) + 7 + msgp.StringPrefixSize + len(z.APIKey) + 10 + msgp.StringPrefixSize + len(z.MasterKey)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Hook) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxpk uint32
	zxpk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxpk > 0 {
		zxpk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "async":
			z.Async, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "event":
			z.Event, err = dc.ReadString()
			if err != nil {
				return
			}
		case "url":
			z.URL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "timeout":
			z.Timeout, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Hook) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "async"
	err = en.Append(0x84, 0xa5, 0x61, 0x73, 0x79, 0x6e, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Async)
	if err != nil {
		return
	}
	// write "event"
	err = en.Append(0xa5, 0x65, 0x76, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Event)
	if err != nil {
		return
	}
	// write "url"
	err = en.Append(0xa3, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.URL)
	if err != nil {
		return
	}
	// write "timeout"
	err = en.Append(0xa7, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Timeout)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Hook) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "async"
	o = append(o, 0x84, 0xa5, 0x61, 0x73, 0x79, 0x6e, 0x63)
	o = msgp.AppendBool(o, z.Async)
	// string "event"
	o = append(o, 0xa5, 0x65, 0x76, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Event)
	// string "url"
	o = append(o, 0xa3, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.URL)
	// string "timeout"
	o = append(o, 0xa7, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74)
	o = msgp.AppendInt(o, z.Timeout)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Hook) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zdnj uint32
	zdnj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zdnj > 0 {
		zdnj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "async":
			z.Async, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "event":
			z.Event, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "url":
			z.URL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "timeout":
			z.Timeout, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Hook) Msgsize() (s int) {
	s = 1 + 6 + msgp.BoolSize + 6 + msgp.StringPrefixSize + len(z.Event) + 4 + msgp.StringPrefixSize + len(z.URL) + 8 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NewAuthConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsnv uint32
	zsnv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsnv > 0 {
		zsnv--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "login_id_keys":
			var zkgt uint32
			zkgt, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.LoginIDKeys) >= int(zkgt) {
				z.LoginIDKeys = (z.LoginIDKeys)[:zkgt]
			} else {
				z.LoginIDKeys = make([]string, zkgt)
			}
			for zobc := range z.LoginIDKeys {
				z.LoginIDKeys[zobc], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "custom_token_secret":
			z.CustomTokenSecret, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *NewAuthConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "login_id_keys"
	err = en.Append(0x82, 0xad, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.LoginIDKeys)))
	if err != nil {
		return
	}
	for zobc := range z.LoginIDKeys {
		err = en.WriteString(z.LoginIDKeys[zobc])
		if err != nil {
			return
		}
	}
	// write "custom_token_secret"
	err = en.Append(0xb3, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.CustomTokenSecret)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NewAuthConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "login_id_keys"
	o = append(o, 0x82, 0xad, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.LoginIDKeys)))
	for zobc := range z.LoginIDKeys {
		o = msgp.AppendString(o, z.LoginIDKeys[zobc])
	}
	// string "custom_token_secret"
	o = append(o, 0xb3, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	o = msgp.AppendString(o, z.CustomTokenSecret)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NewAuthConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zema uint32
	zema, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zema > 0 {
		zema--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "login_id_keys":
			var zpez uint32
			zpez, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.LoginIDKeys) >= int(zpez) {
				z.LoginIDKeys = (z.LoginIDKeys)[:zpez]
			} else {
				z.LoginIDKeys = make([]string, zpez)
			}
			for zobc := range z.LoginIDKeys {
				z.LoginIDKeys[zobc], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "custom_token_secret":
			z.CustomTokenSecret, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NewAuthConfiguration) Msgsize() (s int) {
	s = 1 + 14 + msgp.ArrayHeaderSize
	for zobc := range z.LoginIDKeys {
		s += msgp.StringPrefixSize + len(z.LoginIDKeys[zobc])
	}
	s += 20 + msgp.StringPrefixSize + len(z.CustomTokenSecret)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NewForgotPasswordConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zqke uint32
	zqke, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zqke > 0 {
		zqke--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "app_name":
			z.AppName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "url_prefix":
			z.URLPrefix, err = dc.ReadString()
			if err != nil {
				return
			}
		case "secure_match":
			z.SecureMatch, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "sender_name":
			z.SenderName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "sender":
			z.Sender, err = dc.ReadString()
			if err != nil {
				return
			}
		case "subject":
			z.Subject, err = dc.ReadString()
			if err != nil {
				return
			}
		case "reply_to_name":
			z.ReplyToName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "reply_to":
			z.ReplyTo, err = dc.ReadString()
			if err != nil {
				return
			}
		case "reset_url_lifetime":
			z.ResetURLLifetime, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "success_redirect":
			z.SuccessRedirect, err = dc.ReadString()
			if err != nil {
				return
			}
		case "error_redirect":
			z.ErrorRedirect, err = dc.ReadString()
			if err != nil {
				return
			}
		case "email_text_url":
			z.EmailTextURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "email_html_url":
			z.EmailHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "reset_html_url":
			z.ResetHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "reset_success_html_url":
			z.ResetSuccessHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "reset_error_html_url":
			z.ResetErrorHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *NewForgotPasswordConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 16
	// write "app_name"
	err = en.Append(0xde, 0x0, 0x10, 0xa8, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.AppName)
	if err != nil {
		return
	}
	// write "url_prefix"
	err = en.Append(0xaa, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteString(z.URLPrefix)
	if err != nil {
		return
	}
	// write "secure_match"
	err = en.Append(0xac, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.SecureMatch)
	if err != nil {
		return
	}
	// write "sender_name"
	err = en.Append(0xab, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SenderName)
	if err != nil {
		return
	}
	// write "sender"
	err = en.Append(0xa6, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Sender)
	if err != nil {
		return
	}
	// write "subject"
	err = en.Append(0xa7, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Subject)
	if err != nil {
		return
	}
	// write "reply_to_name"
	err = en.Append(0xad, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyToName)
	if err != nil {
		return
	}
	// write "reply_to"
	err = en.Append(0xa8, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyTo)
	if err != nil {
		return
	}
	// write "reset_url_lifetime"
	err = en.Append(0xb2, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.ResetURLLifetime)
	if err != nil {
		return
	}
	// write "success_redirect"
	err = en.Append(0xb0, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SuccessRedirect)
	if err != nil {
		return
	}
	// write "error_redirect"
	err = en.Append(0xae, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ErrorRedirect)
	if err != nil {
		return
	}
	// write "email_text_url"
	err = en.Append(0xae, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.EmailTextURL)
	if err != nil {
		return
	}
	// write "email_html_url"
	err = en.Append(0xae, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.EmailHTMLURL)
	if err != nil {
		return
	}
	// write "reset_html_url"
	err = en.Append(0xae, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ResetHTMLURL)
	if err != nil {
		return
	}
	// write "reset_success_html_url"
	err = en.Append(0xb6, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ResetSuccessHTMLURL)
	if err != nil {
		return
	}
	// write "reset_error_html_url"
	err = en.Append(0xb4, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ResetErrorHTMLURL)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NewForgotPasswordConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 16
	// string "app_name"
	o = append(o, 0xde, 0x0, 0x10, 0xa8, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.AppName)
	// string "url_prefix"
	o = append(o, 0xaa, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78)
	o = msgp.AppendString(o, z.URLPrefix)
	// string "secure_match"
	o = append(o, 0xac, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68)
	o = msgp.AppendBool(o, z.SecureMatch)
	// string "sender_name"
	o = append(o, 0xab, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.SenderName)
	// string "sender"
	o = append(o, 0xa6, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72)
	o = msgp.AppendString(o, z.Sender)
	// string "subject"
	o = append(o, 0xa7, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74)
	o = msgp.AppendString(o, z.Subject)
	// string "reply_to_name"
	o = append(o, 0xad, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.ReplyToName)
	// string "reply_to"
	o = append(o, 0xa8, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f)
	o = msgp.AppendString(o, z.ReplyTo)
	// string "reset_url_lifetime"
	o = append(o, 0xb2, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt(o, z.ResetURLLifetime)
	// string "success_redirect"
	o = append(o, 0xb0, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74)
	o = msgp.AppendString(o, z.SuccessRedirect)
	// string "error_redirect"
	o = append(o, 0xae, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74)
	o = msgp.AppendString(o, z.ErrorRedirect)
	// string "email_text_url"
	o = append(o, 0xae, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.EmailTextURL)
	// string "email_html_url"
	o = append(o, 0xae, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.EmailHTMLURL)
	// string "reset_html_url"
	o = append(o, 0xae, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.ResetHTMLURL)
	// string "reset_success_html_url"
	o = append(o, 0xb6, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.ResetSuccessHTMLURL)
	// string "reset_error_html_url"
	o = append(o, 0xb4, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.ResetErrorHTMLURL)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NewForgotPasswordConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zqyh uint32
	zqyh, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zqyh > 0 {
		zqyh--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "app_name":
			z.AppName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "url_prefix":
			z.URLPrefix, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "secure_match":
			z.SecureMatch, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "sender_name":
			z.SenderName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "sender":
			z.Sender, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "subject":
			z.Subject, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "reply_to_name":
			z.ReplyToName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "reply_to":
			z.ReplyTo, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "reset_url_lifetime":
			z.ResetURLLifetime, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "success_redirect":
			z.SuccessRedirect, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "error_redirect":
			z.ErrorRedirect, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "email_text_url":
			z.EmailTextURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "email_html_url":
			z.EmailHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "reset_html_url":
			z.ResetHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "reset_success_html_url":
			z.ResetSuccessHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "reset_error_html_url":
			z.ResetErrorHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NewForgotPasswordConfiguration) Msgsize() (s int) {
	s = 3 + 9 + msgp.StringPrefixSize + len(z.AppName) + 11 + msgp.StringPrefixSize + len(z.URLPrefix) + 13 + msgp.BoolSize + 12 + msgp.StringPrefixSize + len(z.SenderName) + 7 + msgp.StringPrefixSize + len(z.Sender) + 8 + msgp.StringPrefixSize + len(z.Subject) + 14 + msgp.StringPrefixSize + len(z.ReplyToName) + 9 + msgp.StringPrefixSize + len(z.ReplyTo) + 19 + msgp.IntSize + 17 + msgp.StringPrefixSize + len(z.SuccessRedirect) + 15 + msgp.StringPrefixSize + len(z.ErrorRedirect) + 15 + msgp.StringPrefixSize + len(z.EmailTextURL) + 15 + msgp.StringPrefixSize + len(z.EmailHTMLURL) + 15 + msgp.StringPrefixSize + len(z.ResetHTMLURL) + 23 + msgp.StringPrefixSize + len(z.ResetSuccessHTMLURL) + 21 + msgp.StringPrefixSize + len(z.ResetErrorHTMLURL)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NewNexmoConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zyzr uint32
	zyzr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zyzr > 0 {
		zyzr--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "api_key":
			z.APIKey, err = dc.ReadString()
			if err != nil {
				return
			}
		case "secret":
			z.APISecret, err = dc.ReadString()
			if err != nil {
				return
			}
		case "from":
			z.From, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z NewNexmoConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "api_key"
	err = en.Append(0x83, 0xa7, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.APIKey)
	if err != nil {
		return
	}
	// write "secret"
	err = en.Append(0xa6, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.APISecret)
	if err != nil {
		return
	}
	// write "from"
	err = en.Append(0xa4, 0x66, 0x72, 0x6f, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.From)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z NewNexmoConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "api_key"
	o = append(o, 0x83, 0xa7, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.APIKey)
	// string "secret"
	o = append(o, 0xa6, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	o = msgp.AppendString(o, z.APISecret)
	// string "from"
	o = append(o, 0xa4, 0x66, 0x72, 0x6f, 0x6d)
	o = msgp.AppendString(o, z.From)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NewNexmoConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zywj uint32
	zywj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zywj > 0 {
		zywj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "api_key":
			z.APIKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "secret":
			z.APISecret, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "from":
			z.From, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NewNexmoConfiguration) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.APIKey) + 7 + msgp.StringPrefixSize + len(z.APISecret) + 5 + msgp.StringPrefixSize + len(z.From)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NewSMTPConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zjpj uint32
	zjpj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjpj > 0 {
		zjpj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "host":
			z.Host, err = dc.ReadString()
			if err != nil {
				return
			}
		case "port":
			z.Port, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "mode":
			z.Mode, err = dc.ReadString()
			if err != nil {
				return
			}
		case "login":
			z.Login, err = dc.ReadString()
			if err != nil {
				return
			}
		case "password":
			z.Password, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *NewSMTPConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "host"
	err = en.Append(0x85, 0xa4, 0x68, 0x6f, 0x73, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Host)
	if err != nil {
		return
	}
	// write "port"
	err = en.Append(0xa4, 0x70, 0x6f, 0x72, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Port)
	if err != nil {
		return
	}
	// write "mode"
	err = en.Append(0xa4, 0x6d, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Mode)
	if err != nil {
		return
	}
	// write "login"
	err = en.Append(0xa5, 0x6c, 0x6f, 0x67, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Login)
	if err != nil {
		return
	}
	// write "password"
	err = en.Append(0xa8, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Password)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NewSMTPConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "host"
	o = append(o, 0x85, 0xa4, 0x68, 0x6f, 0x73, 0x74)
	o = msgp.AppendString(o, z.Host)
	// string "port"
	o = append(o, 0xa4, 0x70, 0x6f, 0x72, 0x74)
	o = msgp.AppendInt(o, z.Port)
	// string "mode"
	o = append(o, 0xa4, 0x6d, 0x6f, 0x64, 0x65)
	o = msgp.AppendString(o, z.Mode)
	// string "login"
	o = append(o, 0xa5, 0x6c, 0x6f, 0x67, 0x69, 0x6e)
	o = msgp.AppendString(o, z.Login)
	// string "password"
	o = append(o, 0xa8, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64)
	o = msgp.AppendString(o, z.Password)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NewSMTPConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zzpf uint32
	zzpf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zzpf > 0 {
		zzpf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "host":
			z.Host, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "port":
			z.Port, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "mode":
			z.Mode, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "login":
			z.Login, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "password":
			z.Password, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NewSMTPConfiguration) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Host) + 5 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Mode) + 6 + msgp.StringPrefixSize + len(z.Login) + 9 + msgp.StringPrefixSize + len(z.Password)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NewSSOConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zeth uint32
	zeth, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zeth > 0 {
		zeth--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "url_prefix":
			z.URLPrefix, err = dc.ReadString()
			if err != nil {
				return
			}
		case "js_sdk_cdn_url":
			z.JSSDKCDNURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "state_jwt_secret":
			z.StateJWTSecret, err = dc.ReadString()
			if err != nil {
				return
			}
		case "auto_link_provider_keys":
			var zsbz uint32
			zsbz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.AutoLinkProviderKeys) >= int(zsbz) {
				z.AutoLinkProviderKeys = (z.AutoLinkProviderKeys)[:zsbz]
			} else {
				z.AutoLinkProviderKeys = make([]string, zsbz)
			}
			for zrfe := range z.AutoLinkProviderKeys {
				z.AutoLinkProviderKeys[zrfe], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "allowed_callback_urls":
			var zrjx uint32
			zrjx, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.AllowedCallbackURLs) >= int(zrjx) {
				z.AllowedCallbackURLs = (z.AllowedCallbackURLs)[:zrjx]
			} else {
				z.AllowedCallbackURLs = make([]string, zrjx)
			}
			for zgmo := range z.AllowedCallbackURLs {
				z.AllowedCallbackURLs[zgmo], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "providers":
			var zawn uint32
			zawn, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Providers) >= int(zawn) {
				z.Providers = (z.Providers)[:zawn]
			} else {
				z.Providers = make([]SSOProviderConfiguration, zawn)
			}
			for ztaf := range z.Providers {
				err = z.Providers[ztaf].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *NewSSOConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "url_prefix"
	err = en.Append(0x86, 0xaa, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteString(z.URLPrefix)
	if err != nil {
		return
	}
	// write "js_sdk_cdn_url"
	err = en.Append(0xae, 0x6a, 0x73, 0x5f, 0x73, 0x64, 0x6b, 0x5f, 0x63, 0x64, 0x6e, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.JSSDKCDNURL)
	if err != nil {
		return
	}
	// write "state_jwt_secret"
	err = en.Append(0xb0, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6a, 0x77, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.StateJWTSecret)
	if err != nil {
		return
	}
	// write "auto_link_provider_keys"
	err = en.Append(0xb7, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.AutoLinkProviderKeys)))
	if err != nil {
		return
	}
	for zrfe := range z.AutoLinkProviderKeys {
		err = en.WriteString(z.AutoLinkProviderKeys[zrfe])
		if err != nil {
			return
		}
	}
	// write "allowed_callback_urls"
	err = en.Append(0xb5, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.AllowedCallbackURLs)))
	if err != nil {
		return
	}
	for zgmo := range z.AllowedCallbackURLs {
		err = en.WriteString(z.AllowedCallbackURLs[zgmo])
		if err != nil {
			return
		}
	}
	// write "providers"
	err = en.Append(0xa9, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Providers)))
	if err != nil {
		return
	}
	for ztaf := range z.Providers {
		err = z.Providers[ztaf].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NewSSOConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "url_prefix"
	o = append(o, 0x86, 0xaa, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78)
	o = msgp.AppendString(o, z.URLPrefix)
	// string "js_sdk_cdn_url"
	o = append(o, 0xae, 0x6a, 0x73, 0x5f, 0x73, 0x64, 0x6b, 0x5f, 0x63, 0x64, 0x6e, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.JSSDKCDNURL)
	// string "state_jwt_secret"
	o = append(o, 0xb0, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6a, 0x77, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	o = msgp.AppendString(o, z.StateJWTSecret)
	// string "auto_link_provider_keys"
	o = append(o, 0xb7, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.AutoLinkProviderKeys)))
	for zrfe := range z.AutoLinkProviderKeys {
		o = msgp.AppendString(o, z.AutoLinkProviderKeys[zrfe])
	}
	// string "allowed_callback_urls"
	o = append(o, 0xb5, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.AllowedCallbackURLs)))
	for zgmo := range z.AllowedCallbackURLs {
		o = msgp.AppendString(o, z.AllowedCallbackURLs[zgmo])
	}
	// string "providers"
	o = append(o, 0xa9, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Providers)))
	for ztaf := range z.Providers {
		o, err = z.Providers[ztaf].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NewSSOConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zwel uint32
	zwel, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zwel > 0 {
		zwel--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "url_prefix":
			z.URLPrefix, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "js_sdk_cdn_url":
			z.JSSDKCDNURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "state_jwt_secret":
			z.StateJWTSecret, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "auto_link_provider_keys":
			var zrbe uint32
			zrbe, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.AutoLinkProviderKeys) >= int(zrbe) {
				z.AutoLinkProviderKeys = (z.AutoLinkProviderKeys)[:zrbe]
			} else {
				z.AutoLinkProviderKeys = make([]string, zrbe)
			}
			for zrfe := range z.AutoLinkProviderKeys {
				z.AutoLinkProviderKeys[zrfe], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "allowed_callback_urls":
			var zmfd uint32
			zmfd, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.AllowedCallbackURLs) >= int(zmfd) {
				z.AllowedCallbackURLs = (z.AllowedCallbackURLs)[:zmfd]
			} else {
				z.AllowedCallbackURLs = make([]string, zmfd)
			}
			for zgmo := range z.AllowedCallbackURLs {
				z.AllowedCallbackURLs[zgmo], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "providers":
			var zzdc uint32
			zzdc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Providers) >= int(zzdc) {
				z.Providers = (z.Providers)[:zzdc]
			} else {
				z.Providers = make([]SSOProviderConfiguration, zzdc)
			}
			for ztaf := range z.Providers {
				bts, err = z.Providers[ztaf].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NewSSOConfiguration) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.URLPrefix) + 15 + msgp.StringPrefixSize + len(z.JSSDKCDNURL) + 17 + msgp.StringPrefixSize + len(z.StateJWTSecret) + 24 + msgp.ArrayHeaderSize
	for zrfe := range z.AutoLinkProviderKeys {
		s += msgp.StringPrefixSize + len(z.AutoLinkProviderKeys[zrfe])
	}
	s += 22 + msgp.ArrayHeaderSize
	for zgmo := range z.AllowedCallbackURLs {
		s += msgp.StringPrefixSize + len(z.AllowedCallbackURLs[zgmo])
	}
	s += 10 + msgp.ArrayHeaderSize
	for ztaf := range z.Providers {
		s += z.Providers[ztaf].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NewTokenStoreConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zelx uint32
	zelx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zelx > 0 {
		zelx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "secret":
			z.Secret, err = dc.ReadString()
			if err != nil {
				return
			}
		case "expiry":
			z.Expiry, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z NewTokenStoreConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "secret"
	err = en.Append(0x82, 0xa6, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Secret)
	if err != nil {
		return
	}
	// write "expiry"
	err = en.Append(0xa6, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Expiry)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z NewTokenStoreConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "secret"
	o = append(o, 0x82, 0xa6, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	o = msgp.AppendString(o, z.Secret)
	// string "expiry"
	o = append(o, 0xa6, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79)
	o = msgp.AppendInt64(o, z.Expiry)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NewTokenStoreConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbal uint32
	zbal, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbal > 0 {
		zbal--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "secret":
			z.Secret, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "expiry":
			z.Expiry, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NewTokenStoreConfiguration) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.Secret) + 7 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NewTwilioConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zjqz uint32
	zjqz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjqz > 0 {
		zjqz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "account_sid":
			z.AccountSID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "auth_token":
			z.AuthToken, err = dc.ReadString()
			if err != nil {
				return
			}
		case "from":
			z.From, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z NewTwilioConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "account_sid"
	err = en.Append(0x83, 0xab, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.AccountSID)
	if err != nil {
		return
	}
	// write "auth_token"
	err = en.Append(0xaa, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.AuthToken)
	if err != nil {
		return
	}
	// write "from"
	err = en.Append(0xa4, 0x66, 0x72, 0x6f, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.From)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z NewTwilioConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "account_sid"
	o = append(o, 0x83, 0xab, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64)
	o = msgp.AppendString(o, z.AccountSID)
	// string "auth_token"
	o = append(o, 0xaa, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e)
	o = msgp.AppendString(o, z.AuthToken)
	// string "from"
	o = append(o, 0xa4, 0x66, 0x72, 0x6f, 0x6d)
	o = msgp.AppendString(o, z.From)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NewTwilioConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zkct uint32
	zkct, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zkct > 0 {
		zkct--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "account_sid":
			z.AccountSID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "auth_token":
			z.AuthToken, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "from":
			z.From, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NewTwilioConfiguration) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.AccountSID) + 11 + msgp.StringPrefixSize + len(z.AuthToken) + 5 + msgp.StringPrefixSize + len(z.From)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NewUserAuditConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var ztmt uint32
	ztmt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ztmt > 0 {
		ztmt--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "enabled":
			z.Enabled, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "trail_handler_url":
			z.TrailHandlerURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "password":
			err = z.Password.DecodeMsg(dc)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *NewUserAuditConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "enabled"
	err = en.Append(0x83, 0xa7, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Enabled)
	if err != nil {
		return
	}
	// write "trail_handler_url"
	err = en.Append(0xb1, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.TrailHandlerURL)
	if err != nil {
		return
	}
	// write "password"
	err = en.Append(0xa8, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64)
	if err != nil {
		return err
	}
	err = z.Password.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NewUserAuditConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "enabled"
	o = append(o, 0x83, 0xa7, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Enabled)
	// string "trail_handler_url"
	o = append(o, 0xb1, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.TrailHandlerURL)
	// string "password"
	o = append(o, 0xa8, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64)
	o, err = z.Password.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NewUserAuditConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var ztco uint32
	ztco, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztco > 0 {
		ztco--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "enabled":
			z.Enabled, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "trail_handler_url":
			z.TrailHandlerURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "password":
			bts, err = z.Password.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NewUserAuditConfiguration) Msgsize() (s int) {
	s = 1 + 8 + msgp.BoolSize + 18 + msgp.StringPrefixSize + len(z.TrailHandlerURL) + 9 + z.Password.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NewWelcomeEmailConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zana uint32
	zana, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zana > 0 {
		zana--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "enabled":
			z.Enabled, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "url_prefix":
			z.URLPrefix, err = dc.ReadString()
			if err != nil {
				return
			}
		case "sender_name":
			z.SenderName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "sender":
			z.Sender, err = dc.ReadString()
			if err != nil {
				return
			}
		case "subject":
			z.Subject, err = dc.ReadString()
			if err != nil {
				return
			}
		case "reply_to_name":
			z.ReplyToName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "reply_to":
			z.ReplyTo, err = dc.ReadString()
			if err != nil {
				return
			}
		case "text_url":
			z.TextURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "html_url":
			z.HTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *NewWelcomeEmailConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "enabled"
	err = en.Append(0x89, 0xa7, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Enabled)
	if err != nil {
		return
	}
	// write "url_prefix"
	err = en.Append(0xaa, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteString(z.URLPrefix)
	if err != nil {
		return
	}
	// write "sender_name"
	err = en.Append(0xab, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SenderName)
	if err != nil {
		return
	}
	// write "sender"
	err = en.Append(0xa6, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Sender)
	if err != nil {
		return
	}
	// write "subject"
	err = en.Append(0xa7, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Subject)
	if err != nil {
		return
	}
	// write "reply_to_name"
	err = en.Append(0xad, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyToName)
	if err != nil {
		return
	}
	// write "reply_to"
	err = en.Append(0xa8, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyTo)
	if err != nil {
		return
	}
	// write "text_url"
	err = en.Append(0xa8, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.TextURL)
	if err != nil {
		return
	}
	// write "html_url"
	err = en.Append(0xa8, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.HTMLURL)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NewWelcomeEmailConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "enabled"
	o = append(o, 0x89, 0xa7, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Enabled)
	// string "url_prefix"
	o = append(o, 0xaa, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78)
	o = msgp.AppendString(o, z.URLPrefix)
	// string "sender_name"
	o = append(o, 0xab, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.SenderName)
	// string "sender"
	o = append(o, 0xa6, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72)
	o = msgp.AppendString(o, z.Sender)
	// string "subject"
	o = append(o, 0xa7, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74)
	o = msgp.AppendString(o, z.Subject)
	// string "reply_to_name"
	o = append(o, 0xad, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.ReplyToName)
	// string "reply_to"
	o = append(o, 0xa8, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f)
	o = msgp.AppendString(o, z.ReplyTo)
	// string "text_url"
	o = append(o, 0xa8, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.TextURL)
	// string "html_url"
	o = append(o, 0xa8, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.HTMLURL)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NewWelcomeEmailConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var ztyy uint32
	ztyy, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for ztyy > 0 {
		ztyy--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "enabled":
			z.Enabled, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "url_prefix":
			z.URLPrefix, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "sender_name":
			z.SenderName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "sender":
			z.Sender, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "subject":
			z.Subject, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "reply_to_name":
			z.ReplyToName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "reply_to":
			z.ReplyTo, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "text_url":
			z.TextURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "html_url":
			z.HTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NewWelcomeEmailConfiguration) Msgsize() (s int) {
	s = 1 + 8 + msgp.BoolSize + 11 + msgp.StringPrefixSize + len(z.URLPrefix) + 12 + msgp.StringPrefixSize + len(z.SenderName) + 7 + msgp.StringPrefixSize + len(z.Sender) + 8 + msgp.StringPrefixSize + len(z.Subject) + 14 + msgp.StringPrefixSize + len(z.ReplyToName) + 9 + msgp.StringPrefixSize + len(z.ReplyTo) + 9 + msgp.StringPrefixSize + len(z.TextURL) + 9 + msgp.StringPrefixSize + len(z.HTMLURL)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PasswordConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zare uint32
	zare, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zare > 0 {
		zare--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "min_length":
			z.MinLength, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "uppercase_required":
			z.UppercaseRequired, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "lowercase_required":
			z.LowercaseRequired, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "digit_required":
			z.DigitRequired, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "symbol_required":
			z.SymbolRequired, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "minimum_guessable_level":
			z.MinimumGuessableLevel, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "excluded_keywords":
			var zljy uint32
			zljy, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.ExcludedKeywords) >= int(zljy) {
				z.ExcludedKeywords = (z.ExcludedKeywords)[:zljy]
			} else {
				z.ExcludedKeywords = make([]string, zljy)
			}
			for zinl := range z.ExcludedKeywords {
				z.ExcludedKeywords[zinl], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "history_size":
			z.HistorySize, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "history_days":
			z.HistoryDays, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "expiry_days":
			z.ExpiryDays, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PasswordConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "min_length"
	err = en.Append(0x8a, 0xaa, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.MinLength)
	if err != nil {
		return
	}
	// write "uppercase_required"
	err = en.Append(0xb2, 0x75, 0x70, 0x70, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.UppercaseRequired)
	if err != nil {
		return
	}
	// write "lowercase_required"
	err = en.Append(0xb2, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.LowercaseRequired)
	if err != nil {
		return
	}
	// write "digit_required"
	err = en.Append(0xae, 0x64, 0x69, 0x67, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.DigitRequired)
	if err != nil {
		return
	}
	// write "symbol_required"
	err = en.Append(0xaf, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.SymbolRequired)
	if err != nil {
		return
	}
	// write "minimum_guessable_level"
	err = en.Append(0xb7, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.MinimumGuessableLevel)
	if err != nil {
		return
	}
	// write "excluded_keywords"
	err = en.Append(0xb1, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.ExcludedKeywords)))
	if err != nil {
		return
	}
	for zinl := range z.ExcludedKeywords {
		err = en.WriteString(z.ExcludedKeywords[zinl])
		if err != nil {
			return
		}
	}
	// write "history_size"
	err = en.Append(0xac, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.HistorySize)
	if err != nil {
		return
	}
	// write "history_days"
	err = en.Append(0xac, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x79, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.HistoryDays)
	if err != nil {
		return
	}
	// write "expiry_days"
	err = en.Append(0xab, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x79, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.ExpiryDays)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PasswordConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "min_length"
	o = append(o, 0x8a, 0xaa, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68)
	o = msgp.AppendInt(o, z.MinLength)
	// string "uppercase_required"
	o = append(o, 0xb2, 0x75, 0x70, 0x70, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64)
	o = msgp.AppendBool(o, z.UppercaseRequired)
	// string "lowercase_required"
	o = append(o, 0xb2, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64)
	o = msgp.AppendBool(o, z.LowercaseRequired)
	// string "digit_required"
	o = append(o, 0xae, 0x64, 0x69, 0x67, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64)
	o = msgp.AppendBool(o, z.DigitRequired)
	// string "symbol_required"
	o = append(o, 0xaf, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64)
	o = msgp.AppendBool(o, z.SymbolRequired)
	// string "minimum_guessable_level"
	o = append(o, 0xb7, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c)
	o = msgp.AppendInt(o, z.MinimumGuessableLevel)
	// string "excluded_keywords"
	o = append(o, 0xb1, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.ExcludedKeywords)))
	for zinl := range z.ExcludedKeywords {
		o = msgp.AppendString(o, z.ExcludedKeywords[zinl])
	}
	// string "history_size"
	o = append(o, 0xac, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt(o, z.HistorySize)
	// string "history_days"
	o = append(o, 0xac, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x79, 0x73)
	o = msgp.AppendInt(o, z.HistoryDays)
	// string "expiry_days"
	o = append(o, 0xab, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x79, 0x73)
	o = msgp.AppendInt(o, z.ExpiryDays)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PasswordConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zixj uint32
	zixj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zixj > 0 {
		zixj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "min_length":
			z.MinLength, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "uppercase_required":
			z.UppercaseRequired, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "lowercase_required":
			z.LowercaseRequired, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "digit_required":
			z.DigitRequired, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "symbol_required":
			z.SymbolRequired, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "minimum_guessable_level":
			z.MinimumGuessableLevel, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "excluded_keywords":
			var zrsc uint32
			zrsc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.ExcludedKeywords) >= int(zrsc) {
				z.ExcludedKeywords = (z.ExcludedKeywords)[:zrsc]
			} else {
				z.ExcludedKeywords = make([]string, zrsc)
			}
			for zinl := range z.ExcludedKeywords {
				z.ExcludedKeywords[zinl], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "history_size":
			z.HistorySize, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "history_days":
			z.HistoryDays, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "expiry_days":
			z.ExpiryDays, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PasswordConfiguration) Msgsize() (s int) {
	s = 1 + 11 + msgp.IntSize + 19 + msgp.BoolSize + 19 + msgp.BoolSize + 15 + msgp.BoolSize + 16 + msgp.BoolSize + 24 + msgp.IntSize + 18 + msgp.ArrayHeaderSize
	for zinl := range z.ExcludedKeywords {
		s += msgp.StringPrefixSize + len(z.ExcludedKeywords[zinl])
	}
	s += 13 + msgp.IntSize + 13 + msgp.IntSize + 12 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SSOProviderConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zctn uint32
	zctn, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zctn > 0 {
		zctn--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "client_id":
			z.ClientID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "client_secret":
			z.ClientSecret, err = dc.ReadString()
			if err != nil {
				return
			}
		case "scope":
			z.Scope, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SSOProviderConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "name"
	err = en.Append(0x84, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "client_id"
	err = en.Append(0xa9, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ClientID)
	if err != nil {
		return
	}
	// write "client_secret"
	err = en.Append(0xad, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ClientSecret)
	if err != nil {
		return
	}
	// write "scope"
	err = en.Append(0xa5, 0x73, 0x63, 0x6f, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Scope)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SSOProviderConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "name"
	o = append(o, 0x84, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "client_id"
	o = append(o, 0xa9, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64)
	o = msgp.AppendString(o, z.ClientID)
	// string "client_secret"
	o = append(o, 0xad, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	o = msgp.AppendString(o, z.ClientSecret)
	// string "scope"
	o = append(o, 0xa5, 0x73, 0x63, 0x6f, 0x70, 0x65)
	o = msgp.AppendString(o, z.Scope)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SSOProviderConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zswy uint32
	zswy, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zswy > 0 {
		zswy--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "client_id":
			z.ClientID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "client_secret":
			z.ClientSecret, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "scope":
			z.Scope, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SSOProviderConfiguration) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 10 + msgp.StringPrefixSize + len(z.ClientID) + 14 + msgp.StringPrefixSize + len(z.ClientSecret) + 6 + msgp.StringPrefixSize + len(z.Scope)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TenantConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zsvm uint32
	zsvm, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zsvm > 0 {
		zsvm--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "version":
			z.Version, err = dc.ReadString()
			if err != nil {
				return
			}
		case "app_name":
			z.AppName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "app_config":
			err = z.AppConfig.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "user_config":
			err = z.UserConfig.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "hooks":
			var zaoz uint32
			zaoz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Hooks) >= int(zaoz) {
				z.Hooks = (z.Hooks)[:zaoz]
			} else {
				z.Hooks = make([]Hook, zaoz)
			}
			for znsg := range z.Hooks {
				err = z.Hooks[znsg].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "deployment_routes":
			var zfzb uint32
			zfzb, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.DeploymentRoutes) >= int(zfzb) {
				z.DeploymentRoutes = (z.DeploymentRoutes)[:zfzb]
			} else {
				z.DeploymentRoutes = make([]DeploymentRoute, zfzb)
			}
			for zrus := range z.DeploymentRoutes {
				err = z.DeploymentRoutes[zrus].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TenantConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "version"
	err = en.Append(0x86, 0xa7, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Version)
	if err != nil {
		return
	}
	// write "app_name"
	err = en.Append(0xa8, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.AppName)
	if err != nil {
		return
	}
	// write "app_config"
	err = en.Append(0xaa, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67)
	if err != nil {
		return err
	}
	err = z.AppConfig.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "user_config"
	err = en.Append(0xab, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67)
	if err != nil {
		return err
	}
	err = z.UserConfig.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "hooks"
	err = en.Append(0xa5, 0x68, 0x6f, 0x6f, 0x6b, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Hooks)))
	if err != nil {
		return
	}
	for znsg := range z.Hooks {
		err = z.Hooks[znsg].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "deployment_routes"
	err = en.Append(0xb1, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.DeploymentRoutes)))
	if err != nil {
		return
	}
	for zrus := range z.DeploymentRoutes {
		err = z.DeploymentRoutes[zrus].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TenantConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "version"
	o = append(o, 0x86, 0xa7, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Version)
	// string "app_name"
	o = append(o, 0xa8, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.AppName)
	// string "app_config"
	o = append(o, 0xaa, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67)
	o, err = z.AppConfig.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "user_config"
	o = append(o, 0xab, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67)
	o, err = z.UserConfig.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "hooks"
	o = append(o, 0xa5, 0x68, 0x6f, 0x6f, 0x6b, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Hooks)))
	for znsg := range z.Hooks {
		o, err = z.Hooks[znsg].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "deployment_routes"
	o = append(o, 0xb1, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.DeploymentRoutes)))
	for zrus := range z.DeploymentRoutes {
		o, err = z.DeploymentRoutes[zrus].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TenantConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsbo uint32
	zsbo, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsbo > 0 {
		zsbo--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "version":
			z.Version, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "app_name":
			z.AppName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "app_config":
			bts, err = z.AppConfig.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "user_config":
			bts, err = z.UserConfig.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "hooks":
			var zjif uint32
			zjif, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Hooks) >= int(zjif) {
				z.Hooks = (z.Hooks)[:zjif]
			} else {
				z.Hooks = make([]Hook, zjif)
			}
			for znsg := range z.Hooks {
				bts, err = z.Hooks[znsg].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "deployment_routes":
			var zqgz uint32
			zqgz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.DeploymentRoutes) >= int(zqgz) {
				z.DeploymentRoutes = (z.DeploymentRoutes)[:zqgz]
			} else {
				z.DeploymentRoutes = make([]DeploymentRoute, zqgz)
			}
			for zrus := range z.DeploymentRoutes {
				bts, err = z.DeploymentRoutes[zrus].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TenantConfiguration) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Version) + 9 + msgp.StringPrefixSize + len(z.AppName) + 11 + z.AppConfig.Msgsize() + 12 + z.UserConfig.Msgsize() + 6 + msgp.ArrayHeaderSize
	for znsg := range z.Hooks {
		s += z.Hooks[znsg].Msgsize()
	}
	s += 18 + msgp.ArrayHeaderSize
	for zrus := range z.DeploymentRoutes {
		s += z.DeploymentRoutes[zrus].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var ztls uint32
	ztls, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ztls > 0 {
		ztls--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "api_key":
			z.APIKey, err = dc.ReadString()
			if err != nil {
				return
			}
		case "master_key":
			z.MasterKey, err = dc.ReadString()
			if err != nil {
				return
			}
		case "url_prefix":
			z.URLPrefix, err = dc.ReadString()
			if err != nil {
				return
			}
		case "cors":
			var zmvo uint32
			zmvo, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zmvo > 0 {
				zmvo--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "origin":
					z.CORS.Origin, err = dc.ReadString()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "auth":
			var zigk uint32
			zigk, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zigk > 0 {
				zigk--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "login_id_keys":
					var zopb uint32
					zopb, err = dc.ReadArrayHeader()
					if err != nil {
						return
					}
					if cap(z.Auth.LoginIDKeys) >= int(zopb) {
						z.Auth.LoginIDKeys = (z.Auth.LoginIDKeys)[:zopb]
					} else {
						z.Auth.LoginIDKeys = make([]string, zopb)
					}
					for zsnw := range z.Auth.LoginIDKeys {
						z.Auth.LoginIDKeys[zsnw], err = dc.ReadString()
						if err != nil {
							return
						}
					}
				case "custom_token_secret":
					z.Auth.CustomTokenSecret, err = dc.ReadString()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "token_store":
			var zuop uint32
			zuop, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zuop > 0 {
				zuop--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "secret":
					z.TokenStore.Secret, err = dc.ReadString()
					if err != nil {
						return
					}
				case "expiry":
					z.TokenStore.Expiry, err = dc.ReadInt64()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "user_audit":
			var zedl uint32
			zedl, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zedl > 0 {
				zedl--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "enabled":
					z.UserAudit.Enabled, err = dc.ReadBool()
					if err != nil {
						return
					}
				case "trail_handler_url":
					z.UserAudit.TrailHandlerURL, err = dc.ReadString()
					if err != nil {
						return
					}
				case "password":
					err = z.UserAudit.Password.DecodeMsg(dc)
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "forgot_password":
			err = z.ForgotPassword.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "welcome_email":
			err = z.WelcomeEmail.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "sso":
			err = z.SSO.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "user_verification":
			err = z.UserVerification.DecodeMsg(dc)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *UserConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 11
	// write "api_key"
	err = en.Append(0x8b, 0xa7, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.APIKey)
	if err != nil {
		return
	}
	// write "master_key"
	err = en.Append(0xaa, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.MasterKey)
	if err != nil {
		return
	}
	// write "url_prefix"
	err = en.Append(0xaa, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteString(z.URLPrefix)
	if err != nil {
		return
	}
	// write "cors"
	// map header, size 1
	// write "origin"
	err = en.Append(0xa4, 0x63, 0x6f, 0x72, 0x73, 0x81, 0xa6, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.CORS.Origin)
	if err != nil {
		return
	}
	// write "auth"
	// map header, size 2
	// write "login_id_keys"
	err = en.Append(0xa4, 0x61, 0x75, 0x74, 0x68, 0x82, 0xad, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Auth.LoginIDKeys)))
	if err != nil {
		return
	}
	for zsnw := range z.Auth.LoginIDKeys {
		err = en.WriteString(z.Auth.LoginIDKeys[zsnw])
		if err != nil {
			return
		}
	}
	// write "custom_token_secret"
	err = en.Append(0xb3, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Auth.CustomTokenSecret)
	if err != nil {
		return
	}
	// write "token_store"
	// map header, size 2
	// write "secret"
	err = en.Append(0xab, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x82, 0xa6, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.TokenStore.Secret)
	if err != nil {
		return
	}
	// write "expiry"
	err = en.Append(0xa6, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.TokenStore.Expiry)
	if err != nil {
		return
	}
	// write "user_audit"
	// map header, size 3
	// write "enabled"
	err = en.Append(0xaa, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x83, 0xa7, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.UserAudit.Enabled)
	if err != nil {
		return
	}
	// write "trail_handler_url"
	err = en.Append(0xb1, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.UserAudit.TrailHandlerURL)
	if err != nil {
		return
	}
	// write "password"
	err = en.Append(0xa8, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64)
	if err != nil {
		return err
	}
	err = z.UserAudit.Password.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "forgot_password"
	err = en.Append(0xaf, 0x66, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64)
	if err != nil {
		return err
	}
	err = z.ForgotPassword.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "welcome_email"
	err = en.Append(0xad, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c)
	if err != nil {
		return err
	}
	err = z.WelcomeEmail.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "sso"
	err = en.Append(0xa3, 0x73, 0x73, 0x6f)
	if err != nil {
		return err
	}
	err = z.SSO.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "user_verification"
	err = en.Append(0xb1, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = z.UserVerification.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UserConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 11
	// string "api_key"
	o = append(o, 0x8b, 0xa7, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.APIKey)
	// string "master_key"
	o = append(o, 0xaa, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.MasterKey)
	// string "url_prefix"
	o = append(o, 0xaa, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78)
	o = msgp.AppendString(o, z.URLPrefix)
	// string "cors"
	// map header, size 1
	// string "origin"
	o = append(o, 0xa4, 0x63, 0x6f, 0x72, 0x73, 0x81, 0xa6, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e)
	o = msgp.AppendString(o, z.CORS.Origin)
	// string "auth"
	// map header, size 2
	// string "login_id_keys"
	o = append(o, 0xa4, 0x61, 0x75, 0x74, 0x68, 0x82, 0xad, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Auth.LoginIDKeys)))
	for zsnw := range z.Auth.LoginIDKeys {
		o = msgp.AppendString(o, z.Auth.LoginIDKeys[zsnw])
	}
	// string "custom_token_secret"
	o = append(o, 0xb3, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	o = msgp.AppendString(o, z.Auth.CustomTokenSecret)
	// string "token_store"
	// map header, size 2
	// string "secret"
	o = append(o, 0xab, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x82, 0xa6, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74)
	o = msgp.AppendString(o, z.TokenStore.Secret)
	// string "expiry"
	o = append(o, 0xa6, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79)
	o = msgp.AppendInt64(o, z.TokenStore.Expiry)
	// string "user_audit"
	// map header, size 3
	// string "enabled"
	o = append(o, 0xaa, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x83, 0xa7, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64)
	o = msgp.AppendBool(o, z.UserAudit.Enabled)
	// string "trail_handler_url"
	o = append(o, 0xb1, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.UserAudit.TrailHandlerURL)
	// string "password"
	o = append(o, 0xa8, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64)
	o, err = z.UserAudit.Password.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "forgot_password"
	o = append(o, 0xaf, 0x66, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64)
	o, err = z.ForgotPassword.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "welcome_email"
	o = append(o, 0xad, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c)
	o, err = z.WelcomeEmail.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "sso"
	o = append(o, 0xa3, 0x73, 0x73, 0x6f)
	o, err = z.SSO.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "user_verification"
	o = append(o, 0xb1, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	o, err = z.UserVerification.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zupd uint32
	zupd, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zupd > 0 {
		zupd--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "api_key":
			z.APIKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "master_key":
			z.MasterKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "url_prefix":
			z.URLPrefix, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "cors":
			var zome uint32
			zome, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zome > 0 {
				zome--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "origin":
					z.CORS.Origin, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "auth":
			var zrvj uint32
			zrvj, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zrvj > 0 {
				zrvj--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "login_id_keys":
					var zarz uint32
					zarz, bts, err = msgp.ReadArrayHeaderBytes(bts)
					if err != nil {
						return
					}
					if cap(z.Auth.LoginIDKeys) >= int(zarz) {
						z.Auth.LoginIDKeys = (z.Auth.LoginIDKeys)[:zarz]
					} else {
						z.Auth.LoginIDKeys = make([]string, zarz)
					}
					for zsnw := range z.Auth.LoginIDKeys {
						z.Auth.LoginIDKeys[zsnw], bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					}
				case "custom_token_secret":
					z.Auth.CustomTokenSecret, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "token_store":
			var zknt uint32
			zknt, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zknt > 0 {
				zknt--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "secret":
					z.TokenStore.Secret, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "expiry":
					z.TokenStore.Expiry, bts, err = msgp.ReadInt64Bytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "user_audit":
			var zxye uint32
			zxye, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zxye > 0 {
				zxye--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "enabled":
					z.UserAudit.Enabled, bts, err = msgp.ReadBoolBytes(bts)
					if err != nil {
						return
					}
				case "trail_handler_url":
					z.UserAudit.TrailHandlerURL, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "password":
					bts, err = z.UserAudit.Password.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "forgot_password":
			bts, err = z.ForgotPassword.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "welcome_email":
			bts, err = z.WelcomeEmail.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "sso":
			bts, err = z.SSO.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "user_verification":
			bts, err = z.UserVerification.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UserConfiguration) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.APIKey) + 11 + msgp.StringPrefixSize + len(z.MasterKey) + 11 + msgp.StringPrefixSize + len(z.URLPrefix) + 5 + 1 + 7 + msgp.StringPrefixSize + len(z.CORS.Origin) + 5 + 1 + 14 + msgp.ArrayHeaderSize
	for zsnw := range z.Auth.LoginIDKeys {
		s += msgp.StringPrefixSize + len(z.Auth.LoginIDKeys[zsnw])
	}
	s += 20 + msgp.StringPrefixSize + len(z.Auth.CustomTokenSecret) + 12 + 1 + 7 + msgp.StringPrefixSize + len(z.TokenStore.Secret) + 7 + msgp.Int64Size + 11 + 1 + 8 + msgp.BoolSize + 18 + msgp.StringPrefixSize + len(z.UserAudit.TrailHandlerURL) + 9 + z.UserAudit.Password.Msgsize() + 16 + z.ForgotPassword.Msgsize() + 14 + z.WelcomeEmail.Msgsize() + 4 + z.SSO.Msgsize() + 18 + z.UserVerification.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserVerificationConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zlsx uint32
	zlsx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zlsx > 0 {
		zlsx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "url_prefix":
			z.URLPrefix, err = dc.ReadString()
			if err != nil {
				return
			}
		case "auto_update":
			z.AutoUpdate, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "auto_send_on_signup":
			z.AutoSendOnSignup, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "auto_send_on_update":
			z.AutoSendOnUpdate, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "required":
			z.Required, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "criteria":
			z.Criteria, err = dc.ReadString()
			if err != nil {
				return
			}
		case "error_redirect":
			z.ErrorRedirect, err = dc.ReadString()
			if err != nil {
				return
			}
		case "error_html_url":
			z.ErrorHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "keys":
			var zbgy uint32
			zbgy, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Keys) >= int(zbgy) {
				z.Keys = (z.Keys)[:zbgy]
			} else {
				z.Keys = make([]UserVerificationKeyConfiguration, zbgy)
			}
			for zucw := range z.Keys {
				err = z.Keys[zucw].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *UserVerificationConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "url_prefix"
	err = en.Append(0x89, 0xaa, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteString(z.URLPrefix)
	if err != nil {
		return
	}
	// write "auto_update"
	err = en.Append(0xab, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.AutoUpdate)
	if err != nil {
		return
	}
	// write "auto_send_on_signup"
	err = en.Append(0xb3, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.AutoSendOnSignup)
	if err != nil {
		return
	}
	// write "auto_send_on_update"
	err = en.Append(0xb3, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.AutoSendOnUpdate)
	if err != nil {
		return
	}
	// write "required"
	err = en.Append(0xa8, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Required)
	if err != nil {
		return
	}
	// write "criteria"
	err = en.Append(0xa8, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Criteria)
	if err != nil {
		return
	}
	// write "error_redirect"
	err = en.Append(0xae, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ErrorRedirect)
	if err != nil {
		return
	}
	// write "error_html_url"
	err = en.Append(0xae, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ErrorHTMLURL)
	if err != nil {
		return
	}
	// write "keys"
	err = en.Append(0xa4, 0x6b, 0x65, 0x79, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Keys)))
	if err != nil {
		return
	}
	for zucw := range z.Keys {
		err = z.Keys[zucw].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UserVerificationConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "url_prefix"
	o = append(o, 0x89, 0xaa, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78)
	o = msgp.AppendString(o, z.URLPrefix)
	// string "auto_update"
	o = append(o, 0xab, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65)
	o = msgp.AppendBool(o, z.AutoUpdate)
	// string "auto_send_on_signup"
	o = append(o, 0xb3, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70)
	o = msgp.AppendBool(o, z.AutoSendOnSignup)
	// string "auto_send_on_update"
	o = append(o, 0xb3, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65)
	o = msgp.AppendBool(o, z.AutoSendOnUpdate)
	// string "required"
	o = append(o, 0xa8, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Required)
	// string "criteria"
	o = append(o, 0xa8, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61)
	o = msgp.AppendString(o, z.Criteria)
	// string "error_redirect"
	o = append(o, 0xae, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74)
	o = msgp.AppendString(o, z.ErrorRedirect)
	// string "error_html_url"
	o = append(o, 0xae, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.ErrorHTMLURL)
	// string "keys"
	o = append(o, 0xa4, 0x6b, 0x65, 0x79, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Keys)))
	for zucw := range z.Keys {
		o, err = z.Keys[zucw].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserVerificationConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrao uint32
	zrao, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrao > 0 {
		zrao--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "url_prefix":
			z.URLPrefix, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "auto_update":
			z.AutoUpdate, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "auto_send_on_signup":
			z.AutoSendOnSignup, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "auto_send_on_update":
			z.AutoSendOnUpdate, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "required":
			z.Required, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "criteria":
			z.Criteria, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "error_redirect":
			z.ErrorRedirect, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "error_html_url":
			z.ErrorHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "keys":
			var zmbt uint32
			zmbt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Keys) >= int(zmbt) {
				z.Keys = (z.Keys)[:zmbt]
			} else {
				z.Keys = make([]UserVerificationKeyConfiguration, zmbt)
			}
			for zucw := range z.Keys {
				bts, err = z.Keys[zucw].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UserVerificationConfiguration) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.URLPrefix) + 12 + msgp.BoolSize + 20 + msgp.BoolSize + 20 + msgp.BoolSize + 9 + msgp.BoolSize + 9 + msgp.StringPrefixSize + len(z.Criteria) + 15 + msgp.StringPrefixSize + len(z.ErrorRedirect) + 15 + msgp.StringPrefixSize + len(z.ErrorHTMLURL) + 5 + msgp.ArrayHeaderSize
	for zucw := range z.Keys {
		s += z.Keys[zucw].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserVerificationKeyConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zvls uint32
	zvls, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zvls > 0 {
		zvls--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, err = dc.ReadString()
			if err != nil {
				return
			}
		case "code_format":
			z.CodeFormat, err = dc.ReadString()
			if err != nil {
				return
			}
		case "expiry":
			z.Expiry, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "success_redirect":
			z.SuccessRedirect, err = dc.ReadString()
			if err != nil {
				return
			}
		case "success_html_url":
			z.SuccessHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "error_redirect":
			z.ErrorRedirect, err = dc.ReadString()
			if err != nil {
				return
			}
		case "error_html_url":
			z.ErrorHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "provider":
			z.Provider, err = dc.ReadString()
			if err != nil {
				return
			}
		case "provider_config":
			err = z.ProviderConfig.DecodeMsg(dc)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *UserVerificationKeyConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "key"
	err = en.Append(0x89, 0xa3, 0x6b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Key)
	if err != nil {
		return
	}
	// write "code_format"
	err = en.Append(0xab, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.CodeFormat)
	if err != nil {
		return
	}
	// write "expiry"
	err = en.Append(0xa6, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Expiry)
	if err != nil {
		return
	}
	// write "success_redirect"
	err = en.Append(0xb0, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SuccessRedirect)
	if err != nil {
		return
	}
	// write "success_html_url"
	err = en.Append(0xb0, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SuccessHTMLURL)
	if err != nil {
		return
	}
	// write "error_redirect"
	err = en.Append(0xae, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ErrorRedirect)
	if err != nil {
		return
	}
	// write "error_html_url"
	err = en.Append(0xae, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ErrorHTMLURL)
	if err != nil {
		return
	}
	// write "provider"
	err = en.Append(0xa8, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Provider)
	if err != nil {
		return
	}
	// write "provider_config"
	err = en.Append(0xaf, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67)
	if err != nil {
		return err
	}
	err = z.ProviderConfig.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UserVerificationKeyConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "key"
	o = append(o, 0x89, 0xa3, 0x6b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	// string "code_format"
	o = append(o, 0xab, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74)
	o = msgp.AppendString(o, z.CodeFormat)
	// string "expiry"
	o = append(o, 0xa6, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79)
	o = msgp.AppendInt64(o, z.Expiry)
	// string "success_redirect"
	o = append(o, 0xb0, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74)
	o = msgp.AppendString(o, z.SuccessRedirect)
	// string "success_html_url"
	o = append(o, 0xb0, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.SuccessHTMLURL)
	// string "error_redirect"
	o = append(o, 0xae, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74)
	o = msgp.AppendString(o, z.ErrorRedirect)
	// string "error_html_url"
	o = append(o, 0xae, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.ErrorHTMLURL)
	// string "provider"
	o = append(o, 0xa8, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72)
	o = msgp.AppendString(o, z.Provider)
	// string "provider_config"
	o = append(o, 0xaf, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67)
	o, err = z.ProviderConfig.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserVerificationKeyConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjfj uint32
	zjfj, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjfj > 0 {
		zjfj--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "code_format":
			z.CodeFormat, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "expiry":
			z.Expiry, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "success_redirect":
			z.SuccessRedirect, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "success_html_url":
			z.SuccessHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "error_redirect":
			z.ErrorRedirect, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "error_html_url":
			z.ErrorHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "provider":
			z.Provider, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "provider_config":
			bts, err = z.ProviderConfig.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UserVerificationKeyConfiguration) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Key) + 12 + msgp.StringPrefixSize + len(z.CodeFormat) + 7 + msgp.Int64Size + 17 + msgp.StringPrefixSize + len(z.SuccessRedirect) + 17 + msgp.StringPrefixSize + len(z.SuccessHTMLURL) + 15 + msgp.StringPrefixSize + len(z.ErrorRedirect) + 15 + msgp.StringPrefixSize + len(z.ErrorHTMLURL) + 9 + msgp.StringPrefixSize + len(z.Provider) + 16 + z.ProviderConfig.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserVerificationProviderConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zzak uint32
	zzak, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zzak > 0 {
		zzak--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "subject":
			z.Subject, err = dc.ReadString()
			if err != nil {
				return
			}
		case "sender":
			z.Sender, err = dc.ReadString()
			if err != nil {
				return
			}
		case "sender_name":
			z.SenderName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "reply_to":
			z.ReplyTo, err = dc.ReadString()
			if err != nil {
				return
			}
		case "reply_to_name":
			z.ReplyToName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "text_url":
			z.TextURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "html_url":
			z.HTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *UserVerificationProviderConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "subject"
	err = en.Append(0x87, 0xa7, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Subject)
	if err != nil {
		return
	}
	// write "sender"
	err = en.Append(0xa6, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Sender)
	if err != nil {
		return
	}
	// write "sender_name"
	err = en.Append(0xab, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SenderName)
	if err != nil {
		return
	}
	// write "reply_to"
	err = en.Append(0xa8, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyTo)
	if err != nil {
		return
	}
	// write "reply_to_name"
	err = en.Append(0xad, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyToName)
	if err != nil {
		return
	}
	// write "text_url"
	err = en.Append(0xa8, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.TextURL)
	if err != nil {
		return
	}
	// write "html_url"
	err = en.Append(0xa8, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.HTMLURL)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UserVerificationProviderConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "subject"
	o = append(o, 0x87, 0xa7, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74)
	o = msgp.AppendString(o, z.Subject)
	// string "sender"
	o = append(o, 0xa6, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72)
	o = msgp.AppendString(o, z.Sender)
	// string "sender_name"
	o = append(o, 0xab, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.SenderName)
	// string "reply_to"
	o = append(o, 0xa8, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f)
	o = msgp.AppendString(o, z.ReplyTo)
	// string "reply_to_name"
	o = append(o, 0xad, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.ReplyToName)
	// string "text_url"
	o = append(o, 0xa8, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.TextURL)
	// string "html_url"
	o = append(o, 0xa8, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x75, 0x72, 0x6c)
	o = msgp.AppendString(o, z.HTMLURL)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserVerificationProviderConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbtz uint32
	zbtz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbtz > 0 {
		zbtz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "subject":
			z.Subject, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "sender":
			z.Sender, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "sender_name":
			z.SenderName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "reply_to":
			z.ReplyTo, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "reply_to_name":
			z.ReplyToName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "text_url":
			z.TextURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "html_url":
			z.HTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UserVerificationProviderConfiguration) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Subject) + 7 + msgp.StringPrefixSize + len(z.Sender) + 12 + msgp.StringPrefixSize + len(z.SenderName) + 9 + msgp.StringPrefixSize + len(z.ReplyTo) + 14 + msgp.StringPrefixSize + len(z.ReplyToName) + 9 + msgp.StringPrefixSize + len(z.TextURL) + 9 + msgp.StringPrefixSize + len(z.HTMLURL)
	return
}
