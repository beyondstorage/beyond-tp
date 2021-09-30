import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';

import '../../../../common/colors.dart';
import '../../../../widgets/button/index.dart';
import '../../../../widgets/button/constants.dart';
import '../controller.dart';

import 'controller.dart';

import 'form.dart';

class CreateIdentity extends StatelessWidget {
  final CreateIdentityController controller =
      Get.put(CreateIdentityController());
  final CreateTaskController taskController = Get.put(CreateTaskController());
  final GlobalKey<FormState> formKey = GlobalKey<FormState>();

  CreateIdentity();

  void returnCreateTask() {
    taskController.isCreatingIdentity(false);
  }

  void onSubmit() {
    final form = formKey.currentState;

    if (!form!.validate()) {
      controller.autoValidateMode(AutovalidateMode.always);
    } else {
      form.save();
      controller.createIdentity().then((value) {
        if (taskController.step.value == 1) {
          taskController.srcCredential(controller.name.value);
        } else {
          taskController.dstCredential(controller.name.value);
        }
        returnCreateTask();
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 360,
      padding: EdgeInsets.only(top: 32, left: 32),
      child: Column(
        children: [
          GestureDetector(
            onTap: () => returnCreateTask(),
            child: MouseRegion(
              cursor: SystemMouseCursors.click,
              child: Row(
                children: [
                  Icon(Icons.reply, size: 16, color: regularFontColor),
                  SizedBox(width: 10),
                  Text(
                    'Return to create task'.tr,
                    style: TextStyle(
                      color: regularFontColor,
                      fontSize: 12,
                      height: 1.67,
                      fontWeight: FontWeight.w400,
                    ),
                  ),
                ],
              ),
            ),
          ),
          SizedBox(height: 24),
          CreateIdentityForm(formKey, onSubmit),
          SizedBox(height: 24),
          Row(
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              Button(
                child: Text("Cancel".tr),
                type: ButtonType.defaults,
                onPressed: () => returnCreateTask(),
              ),
              SizedBox(width: 12),
              Button(
                child: Text("Confirm".tr),
                type: ButtonType.primary,
                onPressed: onSubmit,
              ),
            ],
          ),
        ],
      ),
    );
  }
}
