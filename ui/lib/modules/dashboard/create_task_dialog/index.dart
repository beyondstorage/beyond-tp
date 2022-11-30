import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../../widgets/dialog/index.dart';
import '../../../widgets/button/index.dart';
import '../../../widgets/button/constants.dart';

import 'create_identity/index.dart';
import 'step.dart';
import 'form.dart';

import 'controller.dart';

class CreateTaskDialog extends StatelessWidget {
  final CreateTaskController controller = Get.put(CreateTaskController());
  final GlobalKey<FormState> sourceFormKey = GlobalKey<FormState>();
  final GlobalKey<FormState> targetFormKey = GlobalKey<FormState>();
  final GlobalKey<FormState> createIdentityFormKey = GlobalKey<FormState>();
  final Function getTasks;

  CreateTaskDialog({required name, required this.getTasks}) {
    controller.name(name);
  }

  Widget getFormContent() {
    if (controller.isCreatingIdentity.value) {
      return Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          CreateIdentity(),
        ],
      );
    }

    return Row(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        CreateTaskStep(),
        CreateTaskForm(sourceFormKey, targetFormKey, onSubmit),
      ],
    );
  }

  bool getNextStepStatus() {
    return controller.isEditingName.value ||
        controller.srcType.value.length == 0;
  }

  bool getCompleteCreationStatus() {
    return controller.isEditingName.value ||
        controller.dstType.value.length == 0;
  }

  void nextStep() {
    final sourceForm = sourceFormKey.currentState;

    if (!sourceForm!.validate()) {
      controller.autoValidateMode(AutovalidateMode.always);
      return;
    }

    sourceForm.save();
    controller.step(controller.step.value + 1);
  }

  void onSubmit() {
    final targetForm = targetFormKey.currentState;

    if (!targetForm!.validate()) {
      controller.autoValidateMode(AutovalidateMode.always);
    } else {
      targetForm.save();
      controller.onSubmit(getTasks);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Obx(
      () => CommonDialog(
        title: 'Create task'.tr,
        width: 800,
        content: getFormContent(),
        actions: controller.isCreatingIdentity.value
            ? Container(
                width: 800,
                height: 56,
              )
            : null,
        buttons: [
          Button(
            child: Text("Cancel".tr),
            type: ButtonType.defaults,
            onPressed: controller.closeDialog,
          ),
          SizedBox(width: 12),
          controller.step.value == 2
              ? Button(
                  child: Text("Complete creation".tr),
                  type: ButtonType.primary,
                  disabled: getCompleteCreationStatus(),
                  onPressed: onSubmit,
                )
              : Button(
                  child: Text("Next step".tr),
                  type: ButtonType.primary,
                  disabled: getNextStepStatus(),
                  onPressed: nextStep,
                ),
        ],
        leftButtons: controller.step.value > 1
            ? [
                Button(
                  child: Text("Previous step".tr),
                  type: ButtonType.defaults,
                  onPressed: () => controller.step(controller.step.value - 1),
                )
              ]
            : [],
        onClose: controller.closeDialog,
      ),
    );
  }
}
