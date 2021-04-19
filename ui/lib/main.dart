// @dart=2.9
import 'package:get/get.dart';
import 'package:flutter/material.dart';

import 'theme/theme.dart';
import 'routes/index.dart';
import 'i10n/translation.dart';

import 'modules/page_not_found/index.dart';

import 'configure_nonweb.dart' if (dart.library.html) 'configure_web.dart';

void main() {
  configureApp();
  runApp(App());
}

class App extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      initialRoute: '/',
      getPages: Routes.getPages,
      unknownRoute: GetPage(name: '/404', page: () => PageNotFound()),
      translations: Messages(),
      locale: Locale('en', 'US'),
      fallbackLocale: Locale('en', 'US'),
      theme: themeData,
    );
  }
}