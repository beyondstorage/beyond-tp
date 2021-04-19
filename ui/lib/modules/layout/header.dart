import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../routes/index.dart';
import '../../common/global.dart';
import '../../common/shared_prefs.dart';

class HeaderController extends GetxController {
  RxString userName = "".obs;

  void getUserName() {
    getConfig("username").then((_userName) {
      userName(_userName);
    });
  }

  void logout() {
    clearConfigs().then((res) {
      Get.offAllNamed(Routes.main);
    });
  }
}

class Header extends StatelessWidget implements PreferredSizeWidget {
  @override
  Size get preferredSize => Size.fromHeight(globalHeaderHeight);

  @override
  Widget build(BuildContext context) {
    final HeaderController c = Get.put(HeaderController());
    c.getUserName();

    return Container(
      color: Theme.of(context).appBarTheme.backgroundColor,
      height: globalHeaderHeight,
      padding: EdgeInsets.symmetric(horizontal: 20.0),
      child: Row(
        mainAxisSize: MainAxisSize.max,
        children: [
          // Logo in v1.0
          // Image(
          //   image: AssetImage("images/logo.png"),
          //   height: 32,
          // ),
          SelectableText(
            "Project name".tr,
            style: Theme.of(context).appBarTheme.textTheme!.headline4
          ),
          Expanded(
            child: Text(""),
          ),
          // Obx(() => SelectableText(
          //   c.userName.value,
          //   style: Theme.of(context).appBarTheme.textTheme.headline6),
          // ),
          // Padding(
          //   padding: EdgeInsets.only(left: 20.0),
          //   child: IconButton(
          //     icon: Icon(
          //       Icons.logout, color: Colors.white,
          //     ),
          //     iconSize: 20.0,
          //     tooltip: "Logout".tr,
          //     onPressed: () => c.logout(),
          //   ),
          // ),
        ],
      ),
    );
  }
}

class MobileHeader extends StatelessWidget implements PreferredSizeWidget {
  @override
  Size get preferredSize => Size.fromHeight(globalHeaderHeight);

  @override
  Widget build(BuildContext context) {
    final HeaderController c = Get.put(HeaderController());

    return AppBar(
      title: Text(
        "Project name".tr,
        style: Theme.of(context).appBarTheme.textTheme!.headline4
      ),
      actions: [
        IconButton(
          icon: Icon(Icons.logout, size: 24.0),
          onPressed: () => c.logout(),
        ),
      ]
    );
  }
}
