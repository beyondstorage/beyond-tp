import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:form_validator/form_validator.dart';

import '../../../common/colors.dart';
import '../../../widgets/select/index.dart';
import '../../../widgets/select/model.dart';
import '../../../widgets/radio_group/index.dart';
import '../../../widgets/radio_group/model.dart';

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
            'Library Type'.tr,
            style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              height: 1.67,
              fontWeight: FontWeight.w500,
            ),
          ),
          SizedBox(height: 8),
          Obx(
            () => Select(
              validator: ValidationBuilder()
                  .minLength(1, 'Please Select Library Type')
                  .build(),
              options: [
                SelectOption(
                  label: 'QIngCloud - QingStor',
                  value: 'Qingstor',
                ),
              ],
              value: controller.type.value,
              onChange: controller.type,
            ),
          ),
          SizedBox(height: 22),
          SelectableText(
            'Identity Name'.tr,
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
                .minLength(1, 'Please Enter Identity Name')
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
            'Please Select The Protocol And Fill In The Corresponding Value'.tr,
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
                'Access Key',
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
                      .minLength(1, 'Please Enter Access key')
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
                'Secret Key',
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
                      .minLength(1, 'Please Enter Secret key')
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
            'Please Fill In The Format Of <Protocol> :// <Host> : <Port>'.tr,
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
                      .minLength(1, 'Please Select Protocol')
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
                      .minLength(1, 'Please Enter Host')
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
                      .minLength(1, 'Please Enter Port')
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
