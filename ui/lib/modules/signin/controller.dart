import 'package:flutter/widgets.dart';
import 'package:get/get.dart';

import '../../routes/index.dart';
import '../../common/shared_prefs.dart';

class SigninController extends GetxController {
  RxString server = "".obs;
  RxString userName = "".obs;
  RxString password = "".obs;

  // bool get disabled() {
  //   if (server.value.isEmpty()) {
  //
  //   }
  //   return server.value.isEmpty();
  // }

  Rx<AutovalidateMode> autoValidateMode = AutovalidateMode.disabled.obs;

  void initUser() {
    getKeys().then((Set<String> keys) {
      if (keys.contains("server") && keys.contains("username")) {
        Get.offAllNamed(Routes.main);
      }
    });
  }

  void onSubmit() {
    saveConfig("username", userName.value);
    saveConfig("password", password.value);
    saveConfig("server", server.value).then((result) {
      if (result) {
        Get.offAllNamed(Routes.main);
      }
    });
  }
}