import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:form_validator/form_validator.dart';

import '../../../../common/colors.dart';
import '../controller.dart';

import 'controller.dart';

class IdentityFormField extends GetView<IdentityFormFieldController> {
  final IdentityFormFieldController controller =
      Get.put(IdentityFormFieldController());
  final CreateTaskController taskController = Get.put(CreateTaskController());
  final GlobalKey<FormState> formKey = GlobalKey<FormState>();

  IdentityFormField();

  String get getValue {
    if (taskController.step.value == 1) {
      return taskController.srcCredential.value;
    } else {
      return taskController.dstCredential.value;
    }
  }

  void handleChange(value) {
    if (taskController.step.value == 1) {
      taskController.srcCredential(value);
    } else {
      taskController.dstCredential(value);
    }
  }

  Widget createIdentity() {
    return SizedBox(
      width: 328,
      child: OutlineButton.icon(
        icon: Icon(Icons.add),
        color: Colors.white,
        borderSide: BorderSide(
          color: primaryColor,
          style: BorderStyle.solid,
        ),
        shape: RoundedRectangleBorder(
            side: BorderSide(
              color: primaryColor,
              style: BorderStyle.solid,
            ),
            borderRadius: BorderRadius.all(Radius.circular(4))),
        textColor: primaryColor,
        label: Container(
          height: 32,
          alignment: Alignment.center,
          child: Text(
            'Create identity'.tr,
            style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              fontWeight: FontWeight.w500,
            ),
          ),
        ),
        onPressed: () => taskController.isCreatingIdentity(true),
      ),
    );
  }

  Widget selectIdentity() {
    return DropdownButtonFormField(
      hint: Text('Please choose'.tr),
      style: TextStyle(fontSize: 12, height: 1),
      decoration: InputDecoration(
        border: const OutlineInputBorder(),
        isDense: true,
        contentPadding: EdgeInsets.symmetric(horizontal: 12, vertical: 10),
      ),
      value: getValue != '' ? getValue : null,
      validator:
          ValidationBuilder().minLength(1, 'Please select identity').build(),
      onChanged: handleChange,
      onSaved: handleChange,
      items: [
        ...controller.identities.value.identities.map(
          (identity) => DropdownMenuItem(
            child: Text(
              identity.name.tr,
            ),
            value: identity.name,
          ),
        ),
      ],
    );
  }

  @override
  Widget build(BuildContext context) {
    controller.getIdentities();

    return Column(
      children: [
        Row(
          crossAxisAlignment: CrossAxisAlignment.end,
          children: [
            SelectableText(
              'Identity'.tr,
              style: TextStyle(
                color: regularFontColor,
                fontSize: 12,
                height: 1.67,
                fontWeight: FontWeight.w500,
              ),
            ),
            SizedBox(width: 8),
            SelectableText(
              'Accessed library credential and endpoint'.tr,
              style: TextStyle(
                color: disableFontColor,
                fontSize: 10,
                fontWeight: FontWeight.w400,
              ),
            ),
          ],
        ),
        SizedBox(height: 8),
        Obx(
          () => controller.identities.value.length() == 0
              ? createIdentity()
              : selectIdentity(),
        ),
      ],
    );
  }
}
