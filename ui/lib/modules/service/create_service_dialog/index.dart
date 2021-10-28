import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../../widgets/dialog/index.dart';
import '../../../widgets/button/index.dart';
import '../../../widgets/button/constants.dart';

import 'form.dart';
import 'controller.dart';

class CreateServiceDialog extends StatelessWidget {
  final CreateServiceController controller = Get.put(CreateServiceController());
  final GlobalKey<FormState> formKey = GlobalKey<FormState>();
  final Function getIdentities;

  CreateServiceDialog({required this.getIdentities});

  void onSubmit() {
    final form = formKey.currentState;

    if (!form!.validate()) {
      controller.autoValidateMode(AutovalidateMode.always);
    } else {
      form.save();
      controller.onSubmit(getIdentities);
    }
  }

  @override
  Widget build(BuildContext context) {
    return CommonDialog(
      title: 'Create service'.tr,
      width: 400,
      content: Container(
        width: 400,
        child: Padding(
          padding: EdgeInsets.symmetric(vertical: 24, horizontal: 32),
          child: CreateServiceForm(formKey, onSubmit),
        ),
      ),
      buttons: [
        Button(
          child: Text("Cancel".tr),
          type: ButtonType.defaults,
          onPressed: controller.closeDialog,
        ),
        SizedBox(width: 12),
        Button(
          child: Text("Confirm".tr),
          type: ButtonType.primary,
          onPressed: onSubmit,
        ),
      ],
      onClose: controller.closeDialog,
    );
  }
}
