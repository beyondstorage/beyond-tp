import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:form_validator/form_validator.dart';

import './controller.dart';

class CreateTaskForm extends GetView<CreateTaskController> {
  final GlobalKey<FormState> formKey;
  final Function onSubmit;

  CreateTaskForm(this.formKey, this.onSubmit);

  @override
  Widget build(BuildContext context) {
    return Form(
      key: formKey,
      autovalidateMode: controller.autoValidateMode.value,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Container(
            padding: EdgeInsets.symmetric(horizontal: 30.0),
            margin: EdgeInsets.only(bottom: 32.0),
            child: Row(crossAxisAlignment: CrossAxisAlignment.start, children: [
              SizedBox(
                width: 124,
                child: SelectableText(
                  'Task name'.tr,
                  style: TextStyle(
                    fontSize: 12,
                    height: 2.67,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              ),
              Expanded(
                child: TextFormField(
                  autofocus: true,
                  decoration: InputDecoration(
                    border: const OutlineInputBorder(),
                    contentPadding: EdgeInsets.fromLTRB(16.0, 10.0, 16.0, 10.0),
                  ),
                  maxLines: 1,
                  style: TextStyle(fontSize: 12),
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.text,
                  validator: ValidationBuilder()
                      .minLength(1, 'Please enter name')
                      .build(),
                  onSaved: controller.name,
                ),
              ),
            ]),
          ),
          Container(
            decoration: new BoxDecoration(
              border: Border(
                bottom: BorderSide(
                  style: BorderStyle.solid,
                  color: Color.fromRGBO(228, 235, 241, 1),
                ),
              ),
            ),
          ),
          Padding(
            padding: EdgeInsets.only(top: 16.0, bottom: 32.0),
            child: SelectableText(
              'Origin library set'.tr,
              style: TextStyle(
                fontSize: 12,
                fontWeight: FontWeight.w600,
                height: 1.67,
              ),
            ),
          ),
          Container(
            padding: EdgeInsets.symmetric(horizontal: 30.0),
            margin: EdgeInsets.only(bottom: 24.0),
            child: Row(
              children: [
                SizedBox(
                  width: 124,
                  child: RichText(
                    text: TextSpan(
                      text: 'Origin library type'.tr,
                      style: TextStyle(
                        fontSize: 12,
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
                    hint: Text('Select origin library type'.tr),
                    isExpanded: true,
                    isDense: true,
                    style: TextStyle(fontSize: 12),
                    decoration: InputDecoration(
                      border: const OutlineInputBorder(),
                    ),
                    value: controller.originLibraryType.value != ''
                        ? controller.originLibraryType.value
                        : null,
                    validator: ValidationBuilder()
                        .minLength(1, 'Please select origin library type')
                        .build(),
                    onChanged: controller.originLibraryType,
                    onSaved: controller.originLibraryType,
                    items: [
                      DropdownMenuItem(
                        child: Text('QingStor'.tr),
                        value: 'qingstor',
                      ),
                      DropdownMenuItem(
                        child: Text('FS'.tr),
                        value: 'fs',
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
            child: Row(crossAxisAlignment: CrossAxisAlignment.start, children: [
              SizedBox(
                width: 124,
                child: RichText(
                  text: TextSpan(
                    text: 'Origin library path'.tr,
                    style: TextStyle(
                      fontSize: 12,
                      height: 2.67,
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
                    contentPadding: EdgeInsets.fromLTRB(16.0, 10.0, 16.0, 10.0),
                  ),
                  maxLines: 1,
                  style: TextStyle(fontSize: 12),
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.text,
                  validator: ValidationBuilder()
                      .minLength(1, 'Please enter origin library url')
                      .build(),
                  onSaved: controller.originLibraryPath,
                ),
              ),
            ]),
          ),
          Container(
            decoration: new BoxDecoration(
              border: Border(
                bottom: BorderSide(
                  style: BorderStyle.solid,
                  color: Color.fromRGBO(228, 235, 241, 1),
                ),
              ),
            ),
          ),
          Padding(
            padding: EdgeInsets.only(top: 16.0, bottom: 32.0),
            child: SelectableText(
              'Target library set'.tr,
              style: TextStyle(
                fontSize: 12,
                fontWeight: FontWeight.w600,
                height: 1.67,
              ),
            ),
          ),
          Container(
            padding: EdgeInsets.symmetric(horizontal: 30.0),
            margin: EdgeInsets.only(bottom: 24.0),
            child: Row(
              children: [
                SizedBox(
                  width: 124,
                  child: RichText(
                    text: TextSpan(
                      text: 'Target library type'.tr,
                      style: TextStyle(
                        fontSize: 12,
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
                    hint: Text('Select target library type'.tr),
                    isExpanded: true,
                    isDense: true,
                    style: TextStyle(fontSize: 12),
                    decoration: InputDecoration(
                      border: const OutlineInputBorder(),
                    ),
                    value: controller.targetLibraryType.value != ''
                        ? controller.targetLibraryType.value
                        : null,
                    validator: ValidationBuilder()
                        .minLength(1, 'Please select target library type')
                        .build(),
                    onChanged: controller.targetLibraryType,
                    onSaved: controller.targetLibraryType,
                    items: [
                      DropdownMenuItem(
                        child: Text('QingStor'.tr),
                        value: 'qingstor',
                      ),
                      DropdownMenuItem(
                        child: Text('FS'.tr),
                        value: 'fs',
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
            child: Row(crossAxisAlignment: CrossAxisAlignment.start, children: [
              SizedBox(
                width: 124,
                child: RichText(
                  text: TextSpan(
                    text: 'Target library path'.tr,
                    style: TextStyle(
                      fontSize: 12,
                      height: 2.67,
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
                    contentPadding: EdgeInsets.fromLTRB(16.0, 10.0, 16.0, 10.0),
                  ),
                  maxLines: 1,
                  style: TextStyle(fontSize: 12),
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.text,
                  validator: ValidationBuilder()
                      .minLength(1, 'Please enter target library url')
                      .build(),
                  onSaved: controller.targetLibraryPath,
                ),
              ),
            ]),
          ),
          Container(
            padding: EdgeInsets.symmetric(horizontal: 30.0),
            margin: EdgeInsets.only(bottom: 32.0),
            child: Row(crossAxisAlignment: CrossAxisAlignment.start, children: [
              SizedBox(
                width: 124,
                child: SelectableText(
                  'Bucket name'.tr,
                  style: TextStyle(
                    fontSize: 12,
                    height: 2.67,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              ),
              Expanded(
                child: TextFormField(
                  decoration: InputDecoration(
                    border: const OutlineInputBorder(),
                    contentPadding: EdgeInsets.fromLTRB(16.0, 10.0, 16.0, 10.0),
                  ),
                  maxLines: 1,
                  style: TextStyle(fontSize: 12),
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.text,
                  onSaved: controller.bucketName,
                ),
              ),
            ]),
          ),
          Container(
            padding: EdgeInsets.symmetric(horizontal: 30.0),
            margin: EdgeInsets.only(bottom: 32.0),
            child: Row(crossAxisAlignment: CrossAxisAlignment.start, children: [
              SizedBox(
                width: 124,
                child: SelectableText(
                  'Credential'.tr,
                  style: TextStyle(
                    fontSize: 12,
                    height: 2.67,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              ),
              Expanded(
                child: TextFormField(
                  decoration: InputDecoration(
                    border: const OutlineInputBorder(),
                    contentPadding: EdgeInsets.fromLTRB(16.0, 10.0, 16.0, 10.0),
                  ),
                  maxLines: 1,
                  style: TextStyle(fontSize: 12),
                  textInputAction: TextInputAction.next,
                  keyboardType: TextInputType.text,
                  onSaved: controller.credential,
                ),
              ),
            ]),
          ),
          Container(
            padding: EdgeInsets.symmetric(horizontal: 30.0),
            margin: EdgeInsets.only(bottom: 32.0),
            child: Row(crossAxisAlignment: CrossAxisAlignment.start, children: [
              SizedBox(
                width: 124,
                child: SelectableText(
                  'location'.tr,
                  style: TextStyle(
                    fontSize: 12,
                    height: 2.67,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              ),
              Expanded(
                child: TextFormField(
                  decoration: InputDecoration(
                    border: const OutlineInputBorder(),
                    contentPadding: EdgeInsets.fromLTRB(16.0, 10.0, 16.0, 10.0),
                  ),
                  maxLines: 1,
                  style: TextStyle(fontSize: 12),
                  textInputAction: TextInputAction.done,
                  keyboardType: TextInputType.text,
                  onSaved: controller.location,
                  onEditingComplete: onSubmit,
                ),
              ),
            ]),
          ),
        ],
      ),
    );
  }
}
