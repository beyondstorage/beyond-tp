import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:form_validator/form_validator.dart';

import 'controller.dart';
import '../../widgets/button/index.dart';
import '../../widgets/button/constants.dart';


class SignInForm extends GetView<SigninController> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();

  void onSubmit() {
    final form = _formKey.currentState;

    if (!form!.validate()) {
      controller.autoValidateMode(AutovalidateMode.always);
    } else {
      form.save();
      controller.onSubmit();
    }
  }

  @override
  Widget build(BuildContext context) {
    const sizedBoxSpace = SizedBox(height: 24);

    return Form(
      key: _formKey,
      autovalidateMode: controller.autoValidateMode.value,
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Text("Project name".tr, style: Theme.of(context).textTheme.headline2),
          sizedBoxSpace,
          TextFormField(
            decoration: InputDecoration(
              border: const OutlineInputBorder(),
              labelText: "Server address".tr,
            ),
            maxLines: 1,
            keyboardType: TextInputType.url,
            initialValue: "http://0.0.0.0:7436",
            textInputAction: TextInputAction.next,
            autovalidateMode: AutovalidateMode.onUserInteraction,
            validator: ValidationBuilder().url().build(),
            onSaved: controller.server,
          ),
          sizedBoxSpace,
          TextFormField(
            decoration: InputDecoration(
              border: const OutlineInputBorder(),
              labelText: "Email".tr,
            ),
            maxLines: 1,
            textInputAction: TextInputAction.next,
            keyboardType: TextInputType.emailAddress,
            autovalidateMode: AutovalidateMode.onUserInteraction,
            validator: ValidationBuilder().email().build(),
            onSaved: controller.userName,
          ),
          sizedBoxSpace,
          TextFormField(
            obscureText: true,
            decoration: InputDecoration(
              border: const OutlineInputBorder(),
              labelText: "Password".tr,
            ),
            maxLines: 1,
            autovalidateMode: AutovalidateMode.onUserInteraction,
            validator: ValidationBuilder().minLength(6).build(),
            onSaved: controller.password,
          ),
          sizedBoxSpace,
          Row(
            children: [
              Expanded(
                child: Button(
                  type: ButtonType.primary,
                  child: Text("Submit".tr),
                  onPressed: onSubmit,
                ),
              ),
            ],
          )
        ],
      ),
    );
  }
}