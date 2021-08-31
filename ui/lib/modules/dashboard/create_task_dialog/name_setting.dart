import 'package:flutter/material.dart';
import 'package:get/get.dart';

import '../../../common/colors.dart';
import 'controller.dart';

class NameSetting extends GetView<CreateTaskController> {
  NameSetting();

  void onConfirm(String name) {
    controller.name(name);
    controller.isEditingName(false);
  }

  @override
  Widget build(BuildContext context) {
    TextEditingController c = TextEditingController.fromValue(
      TextEditingValue(
        text: controller.name.value,
        selection: TextSelection.fromPosition(
          TextPosition(
            affinity: TextAffinity.downstream,
            offset: controller.name.value.length,
          ),
        ),
      ),
    );

    return SizedBox(
      height: 100,
      child: Column(
        children: [
          Obx(
            () => Visibility(
              visible: controller.isEditingName.value,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Row(
                    children: [
                      Container(
                        width: 168,
                        height: 32,
                        child: TextField(
                          controller: c,
                          style: TextStyle(
                            color: regularFontColor,
                            fontSize: 12,
                            fontWeight: FontWeight.w400,
                          ),
                          textAlignVertical: TextAlignVertical.bottom,
                          decoration: InputDecoration(
                            isDense: true,
                            border: OutlineInputBorder(
                              borderSide: BorderSide(color: defaultColor),
                              borderRadius: BorderRadius.circular(4),
                            ),
                            hintText: 'Press enter to confirm'.tr,
                          ),
                          onEditingComplete: () => onConfirm(c.text),
                        ),
                      ),
                      SizedBox(width: 10),
                      IconButton(
                        icon: Icon(Icons.done),
                        iconSize: 14,
                        constraints:
                            BoxConstraints(maxWidth: 14, maxHeight: 14),
                        padding: EdgeInsets.all(0),
                        splashRadius: 1.0,
                        onPressed: () => onConfirm(c.text),
                      ),
                      SizedBox(width: 10),
                      IconButton(
                        icon: Icon(Icons.close),
                        iconSize: 14,
                        constraints:
                            BoxConstraints(maxWidth: 14, maxHeight: 14),
                        padding: EdgeInsets.all(0),
                        splashRadius: 1.0,
                        onPressed: () => controller.isEditingName(false),
                      ),
                    ],
                  ),
                  SizedBox(height: 4),
                  SelectableText(
                    '1-50 characters, can contain letters, numbers, underscores'
                        .tr,
                    style: TextStyle(
                      color: disableFontColor,
                      fontSize: 10,
                      height: 1.5,
                      fontWeight: FontWeight.w400,
                    ),
                  ),
                ],
              ),
            ),
          ),
          Obx(
            () => Visibility(
              visible: !controller.isEditingName.value,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Row(
                    children: [
                      SelectableText(
                        controller.name.value,
                        style: TextStyle(
                          color: regularFontColor,
                          fontSize: 12,
                          height: 1.67,
                          fontWeight: FontWeight.w500,
                        ),
                      ),
                      SizedBox(width: 10),
                      IconButton(
                        icon: Icon(IconData(0xe609, fontFamily: 'tpIcon'), size: 12,),
                        iconSize: 14,
                        constraints:
                            BoxConstraints(maxWidth: 14, maxHeight: 14),
                        padding: EdgeInsets.only(top: 2),
                        splashRadius: 1.0,
                        onPressed: () => controller.isEditingName(true),
                      ),
                    ],
                  ),
                  SizedBox(height: 4),
                  SelectableText(
                    'You can Modify the task name'.tr,
                    style: TextStyle(
                      color: disableFontColor,
                      fontSize: 10,
                      height: 1.5,
                      fontWeight: FontWeight.w400,
                    ),
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
