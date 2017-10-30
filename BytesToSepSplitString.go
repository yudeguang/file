package file

// import (
// 	"log"
// 	"strconv"
// 	"strings"
// )

// //把byte类型数据，根据definition定义的数据结构，转换成以sep字符分隔的文本。
// func BytesToSepSplitString(data []byte, definition []string, sep string) string {
// 	if sep == "" {
// 		log.Fatal("sep分隔符不能为空")
// 	}
// 	arr := []string{}
// 	for _, v := range definition {
// 		//前面两位表示顺序，仅为了便于观察，这里并不需要处理这两位
// 		//第三位中，0表示不读取,s表示读取STRING，I表示读取数字类型
// 		str := ""
// 		dataType := v[2:3]
// 		dataLen, _ := strconv.Atoi(v[3:])
// 		switch dataType {
// 		case "0":
// 			_, data = ReadString(data, dataLen)
// 		case "s":
// 			str, data = ReadString(data, dataLen)
// 		case "i":
// 			switch dataLen {
// 			case 1:
// 				temp := uint8(0)
// 				temp, data = ReadUINT8(data)
// 				str = strconv.Itoa(int(temp))
// 			case 2:
// 				temp := uint16(0)
// 				temp, data = ReadUINT16(data)
// 				str = strconv.Itoa(int(temp))
// 			case 4:
// 				temp := uint32(0)
// 				temp, data = ReadUINT32(data)
// 				str = strconv.Itoa(int(temp))
// 			case 8:
// 				temp := uint64(0)
// 				temp, data = ReadUINT64(data)
// 				str = strconv.Itoa(int(temp))
// 			default:
// 				log.Fatal("definition中i的数据类型长度范围必须是1,2,4,8中的其中一种,而当前长度为:", dataLen)
// 			}

// 		default:
// 			log.Fatal("definition中i的数据类型必须是s,i,0其中的一种，而当前长度为:", dataType)
// 		}
// 		if strings.Contains(str, sep) {
// 			log.Fatal("sep分隔符的选择不符合规范，与目标数据有冲突")
// 		}
// 		arr = append(arr, str)
// 	}
// 	//上述过程处理完，如果data中还有数据，那么说明定义有问题
// 	if len(data) != 0 {
// 		log.Fatal("definition定义有误，请检察")
// 	}
// 	return strings.Join(arr, sep)
// }

// //处理时增加了去除末尾的空格
// func BytesToSepSplitStringTrimRight(data []byte, definition []string, sep string) string {
// 	if sep == "" {
// 		log.Fatal("sep分隔符不能为空")
// 	}
// 	arr := []string{}
// 	for _, v := range definition {
// 		//前面两位表示顺序，仅为了便于观察，这里并不需要处理这两位
// 		//第三位中，0表示不读取,s表示读取STRING，I表示读取数字类型
// 		str := ""
// 		dataType := v[2:3]
// 		dataLen, _ := strconv.Atoi(v[3:])
// 		switch dataType {
// 		case "0":
// 			_, data = ReadString(data, dataLen)
// 		case "s":
// 			str, data = ReadString(data, dataLen)
// 		case "i":
// 			switch dataLen {
// 			case 1:
// 				temp := uint8(0)
// 				temp, data = ReadUINT8(data)
// 				str = strconv.Itoa(int(temp))
// 			case 2:
// 				temp := uint16(0)
// 				temp, data = ReadUINT16(data)
// 				str = strconv.Itoa(int(temp))
// 			case 4:
// 				temp := uint32(0)
// 				temp, data = ReadUINT32(data)
// 				str = strconv.Itoa(int(temp))
// 			case 8:
// 				temp := uint64(0)
// 				temp, data = ReadUINT64(data)
// 				str = strconv.Itoa(int(temp))
// 			default:
// 				log.Fatal("definition中i的数据类型长度范围必须是1,2,4,8中的其中一种,而当前长度为:", dataLen)
// 			}

// 		default:
// 			log.Fatal("definition中i的数据类型必须是s,i,0其中的一种，而当前长度为:", dataType)
// 		}
// 		if strings.Contains(str, sep) {
// 			log.Fatal("sep分隔符的选择不符合规范，与目标数据有冲突")
// 		}
// 		str = strings.TrimRight(str, " ")
// 		arr = append(arr, str)
// 	}
// 	//上述过程处理完，如果data中还有数据，那么说明定义有问题
// 	if len(data) != 0 {
// 		log.Fatal("definition定义有误，请检察")
// 	}
// 	return strings.Join(arr, sep)
// }

// /*
// definition定义方法:definition := []string{"01s69", "02i2", "0304"}。其中:
// 1.前两位,01,02,03,表示序列，防止definition定义得太长时看不清楚前后关系。
// 2.第3位，s,i,0表示数据类型,s表示string类型,i表示数字类型,0表示后续数据直接丢弃。
// 3.第4位及之后，表示数据长度。
// 例：
// func example() {
// 	data := []byte{0x33, 0x53, 0x55, 0x47, 0x4D, 0x34, 0x41, 0x37, 0x30, 0x43,
// 		0x47, 0x30, 0x30, 0x30, 0x32, 0x31, 0x37, 0x5F, 0xAA, 0x00, 0x00, 0x32,
// 		0x30, 0x31, 0x31, 0x30, 0x39, 0x30, 0x31, 0x00, 0x00, 0x00}
// 	definition := []string{"01s17", "02i4", "03s8", "0403"}

// 	s := BytesToSepSplitString(data, definition, "@")
// 	log.Println(s)
// }

// */
