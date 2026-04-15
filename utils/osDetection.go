package utils

import (

)

func DetectOs(ttlValue uint8) string {
    if ttlValue <= 0 {
        return "Invalid TTL"
    } else if ttlValue <= 64 {
        return "Linux/Mac/Android"
    } else if ttlValue <= 128 {
        return "Windows"
    } else if ttlValue <= 255 {
        return "Network Device (Cisco/Solaris)"
    }
    return "Unknown"
}