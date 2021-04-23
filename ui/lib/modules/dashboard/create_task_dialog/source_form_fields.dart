import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:form_validator/form_validator.dart';

import './controller.dart';

class SourceFormFields extends GetView<CreateTaskController> {
  FocusNode srcFocusNode = FocusNode();

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Padding(
          padding: EdgeInsets.only(top: 16.0, bottom: 32.0),
          child: SelectableText(
            'Source set'.tr,
            style: TextStyle(
              fontSize: 12,
              height: 1.67,
              color: Colors.black,
              fontWeight: FontWeight.w600,
            ),
          ),
        ),
        Container(
          padding: EdgeInsets.symmetric(horizontal: 30.0),
          margin: EdgeInsets.only(bottom: 24.0),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Container(
                width: 124,
                height: 32,
                alignment: Alignment.centerLeft,
                child: RichText(
                  text: TextSpan(
                    text: 'Source type'.tr,
                    style: TextStyle(
                      fontSize: 12,
                      color: Theme.of(context).primaryColorLight,
                      fontWeight: FontWeight.w500,
                    ),
                    children: [
                      TextSpan(
                        text: '*',
                        style: TextStyle(
                          color: Colors.red,
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              Expanded(
                child: DropdownButtonFormField(
                  hint: Text('Select source type'.tr),
                  style: TextStyle(fontSize: 12, height: 1),
                  decoration: InputDecoration(
                    border: const OutlineInputBorder(),
                    isDense: true,
                    contentPadding:
                        EdgeInsets.symmetric(horizontal: 12, vertical: 10),
                  ),
                  value: controller.srcType.value != ''
                      ? controller.srcType.value
                      : null,
                  validator: ValidationBuilder()
                      .minLength(1, 'Please select source type')
                      .build(),
                  onChanged: (String? value) {
                    controller.srcType(value);
                    srcFocusNode.requestFocus();
                  },
                  onSaved: controller.srcType,
                  items: [
                    DropdownMenuItem(
                      child: Text(
                        'QingStor'.tr,
                        style: TextStyle(
                          color: Theme.of(context).primaryColorLight,
                        ),
                      ),
                      value: 'Qingstor',
                    ),
                    DropdownMenuItem(
                      child: Text(
                        'FS'.tr,
                        style: TextStyle(
                          color: Theme.of(context).primaryColorLight,
                        ),
                      ),
                      value: 'Fs',
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
        Container(
          padding: EdgeInsets.symmetric(horizontal: 30.0),
          margin: EdgeInsets.only(bottom: 32.0),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Container(
                width: 124,
                height: 32,
                alignment: Alignment.centerLeft,
                child: Text(
                  'Source path'.tr,
                  style: TextStyle(
                    fontSize: 12,
                    color: Theme.of(context).primaryColorLight,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              ),
              Expanded(
                child: TextFormField(
                  decoration: InputDecoration(
                    border: const OutlineInputBorder(),
                    isDense: true,
                    contentPadding:
                        EdgeInsets.symmetric(horizontal: 12, vertical: 14),
                  ),
                  focusNode: srcFocusNode,
                  maxLines: 1,
                  style: TextStyle(fontSize: 12),
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.text,
                  initialValue: '/',
                  onSaved: controller.srcPath,
                ),
              ),
            ],
          ),
        ),
        Obx(
          () => Visibility(
            visible: controller.srcType.value == 'Qingstor',
            child: Container(
              padding: EdgeInsets.symmetric(horizontal: 30.0),
              margin: EdgeInsets.only(bottom: 32.0),
              child: Row(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Container(
                    width: 124,
                    height: 32,
                    alignment: Alignment.centerLeft,
                    child: RichText(
                      text: TextSpan(
                        text: 'Credential'.tr,
                        style: TextStyle(
                          fontSize: 12,
                          color: Theme.of(context).primaryColorLight,
                          fontWeight: FontWeight.w500,
                        ),
                        children: [
                          TextSpan(
                            text: '*',
                            style: TextStyle(
                              color: Colors.red,
                            ),
                          ),
                        ],
                      ),
                    ),
                  ),
                  Expanded(
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
                          .minLength(1, 'Please enter credential')
                          .build(),
                      onSaved: controller.srcCredential,
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
        Obx(
          () => Visibility(
            visible: controller.srcType.value == 'Qingstor',
            child: Container(
              padding: EdgeInsets.symmetric(horizontal: 30.0),
              margin: EdgeInsets.only(bottom: 32.0),
              child: Row(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Container(
                    width: 124,
                    height: 32,
                    alignment: Alignment.centerLeft,
                    child: Text(
                      'Endpoint'.tr,
                      style: TextStyle(
                        fontSize: 12,
                        color: Theme.of(context).primaryColorLight,
                        fontWeight: FontWeight.w500,
                      ),
                    ),
                  ),
                  Expanded(
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
                      initialValue: 'https:qingstor.com',
                      onSaved: controller.srcEndpoint,
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
        Obx(
          () => Visibility(
            visible: controller.srcType.value == 'Qingstor',
            child: Container(
              padding: EdgeInsets.symmetric(horizontal: 30.0),
              margin: EdgeInsets.only(bottom: 32.0),
              child: Row(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Container(
                    width: 124,
                    height: 32,
                    alignment: Alignment.centerLeft,
                    child: RichText(
                      text: TextSpan(
                        text: 'Bucket name'.tr,
                        style: TextStyle(
                          fontSize: 12,
                          color: Theme.of(context).primaryColorLight,
                          fontWeight: FontWeight.w500,
                        ),
                        children: [
                          TextSpan(
                            text: '*',
                            style: TextStyle(
                              color: Colors.red,
                            ),
                          ),
                        ],
                      ),
                    ),
                  ),
                  Expanded(
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
                          .minLength(1, 'Please enter bucket name')
                          .build(),
                      onSaved: controller.srcBucketName,
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ],
    );
  }
}
