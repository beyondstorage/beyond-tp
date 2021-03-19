import 'package:shared_preferences/shared_preferences.dart';

Future<Set<String>> getKeys() async {
  SharedPreferences prefs = await SharedPreferences.getInstance();

  return prefs.getKeys();
}

Future<String> getConfig(String key) async {
  SharedPreferences prefs = await SharedPreferences.getInstance();

  return prefs.getString(key);
}

Future<bool> saveConfig(String key, String value) async {
  SharedPreferences prefs = await SharedPreferences.getInstance();

  return prefs.setString(key, value);
}

Future<bool> removeConfig(String key) async {
  SharedPreferences prefs = await SharedPreferences.getInstance();

  return prefs.remove(key);
}

Future<bool> clearConfigs() async {
  SharedPreferences prefs = await SharedPreferences.getInstance();

  return prefs.clear();
}
