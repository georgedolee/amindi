package formatter

func symbolFor(code int16) string {
	switch code {
	case 1000:
		return "☀️"
	case 1003:
		return "⛅"
	case 1006, 1009:
		return "☁️"
	case 1030, 1135, 1147:
		return "🌫️"
	case 1063, 1180, 1240:
		return "🌦️"
	case 1150, 1153, 1168, 1171, 1183:
		return "🌧️"
	case 1186, 1189, 1192, 1195, 1243, 1246:
		return "🌧️"
	case 1198, 1201:
		return "🌧️❄️"
	case 1066, 1069, 1204, 1210, 1213:
		return "🌨️"
	case 1207, 1252:
		return "🌨️"
	case 1114, 1117, 1222, 1225, 1255, 1258:
		return "❄️"
	case 1237, 1261, 1264:
		return "🧊"
	case 1273, 1087, 1276:
		return "⛈️"
	case 1279, 1282:
		return "⛈️❄️"
	default:
		return "🌡️"
	}
}
