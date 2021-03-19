import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../widgets/button/index.dart';
import '../../widgets/toolbar/index.dart';

import 'controller.dart';

class Toolbar extends GetView<DashboardController> {
  @override
  Widget build(BuildContext context) {
    return PageToolbar(
      title: 'Tasks'.tr,
      children: [
        Button(
          icon: Icons.add,
          label: Text("New task".tr),
          type: ButtonType.primary,
          onPressed: controller.onPressedNew,
        ),
      ],
    );
  }
}

