package platforms

func AnyAndroidIntentExtras() AndroidIntentExtras {
	return []string{"-e", "any", "android", "-e", "intent", "extras"}
}
