import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:flutter/rendering.dart';

import '../../../common/colors.dart';
import 'name_setting.dart';
import 'controller.dart';

class CreateTaskStep extends GetView<CreateTaskController> {
  CreateTaskStep();

  Widget getStepButton(int step) {
    if (step < controller.step.value) {
      return Container(
        width: 28,
        height: 28,
        alignment: Alignment.center,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(14),
          border: Border.all(color: primaryColor, width: 2),
        ),
        child: Icon(Icons.done, color: primaryColor, size: 14),
      );
    } else if (step == controller.step.value) {
      return Container(
        width: 28,
        height: 28,
        alignment: Alignment.center,
        decoration: BoxDecoration(
          color: primaryColor,
          borderRadius: BorderRadius.circular(14),
        ),
        child: Text(
          step.toString(),
          style: TextStyle(
            color: Colors.white,
            fontSize: 12,
            fontWeight: FontWeight.w500,
          ),
        ),
      );
    } else {
      return Container(
        width: 28,
        height: 28,
        alignment: Alignment.center,
        decoration: BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.circular(14),
          border: Border.all(color: disableFontColor, width: 2),
        ),
        child: Text(
          step.toString(),
          style: TextStyle(
            color: disableFontColor,
            fontSize: 12,
            fontWeight: FontWeight.w500,
          ),
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 260,
      height: 548,
      padding: EdgeInsets.only(top: 32, left: 32),
      decoration: new BoxDecoration(
        border: Border(
          right: BorderSide(
            color: rgba(226, 232, 240, 1),
            width: 1,
          ),
        ),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          NameSetting(),
          GestureDetector(
            onTap: () => controller.step(1),
            child: MouseRegion(
              cursor: SystemMouseCursors.click,
              child: Row(
                children: [
                  Obx(() => getStepButton(1)),
                  SizedBox(width: 12),
                  Text(
                    'Source library settings'.tr,
                    style: TextStyle(
                      color: headlineFontColor,
                      fontSize: 12,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                ],
              ),
            ),
          ),
          SizedBox(height: 8),
          Obx(
            () => Container(
              width: 15,
              height: 60,
              decoration: BoxDecoration(
                border: Border(
                  right: BorderSide(
                    color: controller.step.value == 1
                        ? disableFontColor
                        : primaryColor,
                    width: 2,
                  ),
                ),
              ),
            ),
          ),
          SizedBox(height: 8),
          GestureDetector(
            onTap: () {
              if (controller.srcType.value.length > 0) {
                controller.step(2);
              }
            },
            child: MouseRegion(
              cursor: SystemMouseCursors.click,
              child: Row(
                children: [
                  Obx(() => getStepButton(2)),
                  SizedBox(width: 12),
                  Text(
                    'Target library settings'.tr,
                    style: TextStyle(
                      color: controller.step.value == 1
                          ? disableFontColor
                          : headlineFontColor,
                      fontSize: 12,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                ],
              ),
            ),
          ),
          // SizedBox(height: 8),
          // Container(
          //   width: 15,
          //   height: 60,
          //   decoration: BoxDecoration(
          //       border: Border(
          //     right: BorderSide(
          //       color: controller.step.value == 3
          //           ? primaryColor
          //           : disableFontColor,
          //       width: 2,
          //     ),
          //   )),
          // ),
          // SizedBox(height: 8),
          // Row(
          //   children: [
          //     getStepButton(3),
          //     SizedBox(width: 12),
          //     SelectableText(
          //       'Other Settings'.tr,
          //       style: TextStyle(
          //         color: controller.step.value == 3
          //             ? headlineFontColor
          //             : disableFontColor,
          //         fontSize: 12,
          //         fontWeight: FontWeight.w500,
          //       ),
          //     ),
          //   ],
          // ),
        ],
      ),
    );
  }
}
