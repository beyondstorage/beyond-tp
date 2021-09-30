import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:form_validator/form_validator.dart';

import '../../../../common/colors.dart';
import '../../../../widgets/select/index.dart';
import '../../../../widgets/select/model.dart';
import '../../../../widgets/radio_group/index.dart';
import '../../../../widgets/radio_group/model.dart';

import 'controller.dart';

class CreateIdentityForm extends GetView<CreateIdentityController> {
  final GlobalKey<FormState> formKey;
  final Function onSubmit;

  CreateIdentityForm(this.formKey, this.onSubmit);

  @override
  Widget build(BuildContext context) {
    return Form(
      key: formKey,
      autovalidateMode: controller.autoValidateMode.value,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          SelectableText(
            'Identity name'.tr,
            style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              height: 1.67,
              fontWeight: FontWeight.w500,
            ),
          ),
          SizedBox(height: 8),
          TextFormField(
            decoration: InputDecoration(
              border: const OutlineInputBorder(),
              isDense: true,
              contentPadding:
                  EdgeInsets.symmetric(horizontal: 12, vertical: 14),
            ),
            maxLines: 1,
            style: TextStyle(fontSize: 12),
            textInputAction: TextInputAction.next,
            keyboardType: TextInputType.text,
            validator: ValidationBuilder()
                .minLength(1, 'Please enter identity name')
                .build(),
            onSaved: controller.name,
          ),
          SizedBox(height: 22),
          SelectableText(
            'Credential'.tr,
            style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              height: 1.67,
              fontWeight: FontWeight.w500,
            ),
          ),
          SizedBox(height: 4),
          SelectableText(
            'Please select the protocol and fill in the corresponding value'.tr,
            style: TextStyle(
              color: disableFontColor,
              fontSize: 10,
              height: 1.5,
              fontWeight: FontWeight.w400,
            ),
          ),
          SizedBox(height: 4),
          RadioGroup(
            value: 'hamc',
            options: [
              RadioOption(
                label: 'hamc',
                value: 'hamc',
              ),
            ],
          ),
          SizedBox(height: 8),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              SelectableText(
                'Access key',
                style: TextStyle(
                  color: regularFontColor,
                  fontSize: 12,
                  fontWeight: FontWeight.w400,
                ),
              ),
              SizedBox(
                width: 250,
                child: TextFormField(
                  decoration: InputDecoration(
                    border: const OutlineInputBorder(),
                    isDense: true,
                    contentPadding:
                        EdgeInsets.symmetric(horizontal: 12, vertical: 14),
                  ),
                  maxLines: 1,
                  style: TextStyle(fontSize: 12),
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.text,
                  validator: ValidationBuilder()
                      .minLength(1, 'Please enter access key')
                      .build(),
                  onSaved: controller.credentialAccessKey,
                ),
              ),
            ],
          ),
          SizedBox(height: 16),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              SelectableText(
                'Secret key',
                style: TextStyle(
                  color: regularFontColor,
                  fontSize: 12,
                  fontWeight: FontWeight.w400,
                ),
              ),
              SizedBox(
                width: 250,
                child: TextFormField(
                  decoration: InputDecoration(
                    border: const OutlineInputBorder(),
                    isDense: true,
                    contentPadding:
                        EdgeInsets.symmetric(horizontal: 12, vertical: 14),
                  ),
                  maxLines: 1,
                  style: TextStyle(fontSize: 12),
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.text,
                  validator: ValidationBuilder()
                      .minLength(1, 'Please enter secret key')
                      .build(),
                  onSaved: controller.credentialSecretKey,
                ),
              ),
            ],
          ),
          SizedBox(height: 18),
          SelectableText(
            'Endpoint'.tr,
            style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              height: 1.67,
              fontWeight: FontWeight.w500,
            ),
          ),
          SizedBox(height: 4),
          SelectableText(
            'Please fill in the format of <Protocol> :// <Host> : <Port>'.tr,
            style: TextStyle(
              color: disableFontColor,
              fontSize: 10,
              height: 1.5,
              fontWeight: FontWeight.w400,
            ),
          ),
          SizedBox(height: 8),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              SizedBox(
                width: 76,
                child: Select(
                  options: [
                    SelectOption(label: 'https', value: 'https'),
                    SelectOption(label: 'http', value: 'http'),
                  ],
                  validator: ValidationBuilder()
                      .minLength(1, 'Please select protocol')
                      .build(),
                  onChange: controller.endpointProtocol,
                ),
              ),
              SelectableText(
                '://',
                style: TextStyle(
                  color: regularFontColor,
                  fontSize: 12,
                  fontWeight: FontWeight.w400,
                ),
              ),
              SizedBox(
                width: 160,
                child: TextFormField(
                  decoration: InputDecoration(
                    border: const OutlineInputBorder(),
                    isDense: true,
                    contentPadding:
                        EdgeInsets.symmetric(horizontal: 12, vertical: 14),
                  ),
                  maxLines: 1,
                  style: TextStyle(fontSize: 12),
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.text,
                  validator: ValidationBuilder()
                      .minLength(1, 'Please enter host')
                      .build(),
                  onSaved: controller.endpointHost,
                ),
              ),
              SelectableText(
                ':',
                style: TextStyle(
                  color: regularFontColor,
                  fontSize: 12,
                  fontWeight: FontWeight.w400,
                ),
              ),
              SizedBox(
                width: 52,
                child: TextFormField(
                  decoration: InputDecoration(
                    border: const OutlineInputBorder(),
                    isDense: true,
                    contentPadding:
                        EdgeInsets.symmetric(horizontal: 12, vertical: 14),
                  ),
                  maxLines: 1,
                  style: TextStyle(fontSize: 12),
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.text,
                  validator: ValidationBuilder()
                      .minLength(1, 'Please enter port')
                      .build(),
                  onSaved: controller.endpointPort,
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
