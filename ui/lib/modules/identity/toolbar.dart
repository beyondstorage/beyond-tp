import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../common/colors.dart';
import '../../widgets/button/index.dart';
import '../../widgets/button/constants.dart';

import 'create_identity_dialog/index.dart';
import 'controller.dart';

class Toolbar extends GetView<IdentityController> {
  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Button(
          icon: Icons.add,
          child: Text("Create task".tr),
          type: ButtonType.primary,
          onPressed: () => Get.dialog(
              CreateIdentityDialog(getIdentities: controller.getIdentities)),
        ),
        SizedBox(width: 20),
        Obx(
          () => SelectableText(
            '${controller.identities.value.length()} Identities',
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
