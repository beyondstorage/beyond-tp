import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../common/colors.dart';
import '../../widgets/button/index.dart';
import '../../widgets/button/constants.dart';

import 'create_service_dialog/index.dart';
import 'controller.dart';

class Toolbar extends GetView<ServiceController> {
  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Button(
          icon: Icons.add,
          child: Text("Create service".tr),
          type: ButtonType.primary,
          onPressed: () => Get.dialog(
              CreateServiceDialog(getIdentities: controller.getIdentities)),
        ),
        SizedBox(width: 20),
        Obx(
          () => SelectableText(
            '${controller.identities.value.length()} services',
            style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              fontWeight: FontWeight.w400,
            ),
          ),
        ),
      ],
    );
  }
}
