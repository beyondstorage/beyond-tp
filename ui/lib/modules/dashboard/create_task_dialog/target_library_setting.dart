import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:form_validator/form_validator.dart';

import '../../../common/colors.dart';
import './controller.dart';

class TargetLibrarySetting extends GetView<CreateTaskController> {
  final GlobalKey<FormState> targetFormKey;

  TargetLibrarySetting(this.targetFormKey);

  @override
  Widget build(BuildContext context) {
    return Form(
      key: targetFormKey,
      autovalidateMode: controller.autoValidateMode.value,
      child: Container(
        width: 328,
        padding: EdgeInsets.only(top: 32, left: 32),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            SelectableText(
              'Target Library'.tr,
              style: TextStyle(
                color: regularFontColor,
                fontSize: 12,
                height: 1.67,
                fontWeight: FontWeight.w500,
              ),
            ),
            SizedBox(height: 8),
            DropdownButtonFormField(
              hint: Text('Please Choose'.tr),
              style: TextStyle(fontSize: 12, height: 1),
              decoration: InputDecoration(
                border: const OutlineInputBorder(),
                isDense: true,
                contentPadding:
                    EdgeInsets.symmetric(horizontal: 12, vertical: 10),
              ),
              value: controller.dstType.value != ''
                  ? controller.dstType.value
                  : null,
              validator: ValidationBuilder()
                  .minLength(1, 'Please select source type')
                  .build(),
              onChanged: (String? value) {
                controller.dstType(value);
              },
              onSaved: controller.dstType,
              items: [
                DropdownMenuItem(
                  child: Text(
                    'Qing Cloud - Qingstor'.tr,
                    style: TextStyle(
                      color: Theme.of(context).primaryColorLight,
                    ),
                  ),
                  value: 'Qingstor',
                ),
                DropdownMenuItem(
                  child: Text(
                    'Local Files - FS'.tr,
                    style: TextStyle(
                      color: Theme.of(context).primaryColorLight,
                    ),
                  ),
                  value: 'Fs',
                ),
              ],
            ),
            Obx(
              () => Visibility(
                visible: controller.dstType.value == 'Qingstor',
                child: Column(
                  children: [
                    SizedBox(height: 24),
                    Row(
                      children: [
                        SelectableText(
                          'Bucket Name'.tr,
                          style: TextStyle(
                            color: regularFontColor,
                            fontSize: 12,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                        SizedBox(width: 8),
                        SelectableText(
                          'QingStor Object Storage\'s Bucket Name'.tr,
                          style: TextStyle(
                            color: disableFontColor,
                            fontSize: 10,
                            fontWeight: FontWeight.w400,
                          ),
                        ),
                      ],
                    ),
                    SizedBox(height: 8),
                    TextFormField(
                      decoration: InputDecoration(
                        border: const OutlineInputBorder(),
                        isDense: true,
                        hintText: 'Please Enter, 6 - 63 Characters'.tr,
                        contentPadding:
                            EdgeInsets.symmetric(horizontal: 12, vertical: 14),
                      ),
                      maxLines: 1,
                      style: TextStyle(fontSize: 12),
                      textInputAction: TextInputAction.next,
                      keyboardType: TextInputType.text,
                      onSaved: controller.dstBucketName,
                    ),
                  ],
                ),
              ),
            ),
            Obx(
              () => Visibility(
                visible: controller.dstType.value == 'Qingstor' ||
                    controller.dstType.value == 'Fs',
                child: Column(
                  children: [
                    SizedBox(height: 24),
                    Row(
                      children: [
                        SelectableText(
                          'Work Dir'.tr,
                          style: TextStyle(
                            color: regularFontColor,
                            fontSize: 12,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                        SizedBox(width: 8),
                        SelectableText(
                          'The Current Working Directory For Service'.tr,
                          style: TextStyle(
                            color: disableFontColor,
                            fontSize: 10,
                            fontWeight: FontWeight.w400,
                          ),
                        ),
                      ],
                    ),
                    SizedBox(height: 8),
                    TextFormField(
                      decoration: InputDecoration(
                        border: const OutlineInputBorder(),
                        isDense: true,
                        hintText: 'Please Enter，Default To /'.tr,
                        contentPadding:
                            EdgeInsets.symmetric(horizontal: 12, vertical: 14),
                      ),
                      maxLines: 1,
                      style: TextStyle(fontSize: 12),
                      textInputAction: TextInputAction.next,
                      keyboardType: TextInputType.text,
                      initialValue: '/',
                      onSaved: controller.dstPath,
                    ),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
