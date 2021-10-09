import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:ui/common/colors.dart';
import 'package:ui/modules/dashboard/create_task_dialog/controller.dart';

import '../../../widgets/dialog/index.dart';
import '../../../widgets/button/index.dart';
import '../../../widgets/button/constants.dart';

class IdentityDialog extends StatelessWidget {
  final CreateTaskController controller = Get.put(CreateTaskController());

  @override
  Widget build(BuildContext context) {
    return CommonDialog(
      title: 'QingStor - Services 1'.tr,
      width: 800,
      content: Container(
        width: 320,
        height: 269,
        padding: EdgeInsets.all(30),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                SelectableText(
                  'Credential',
                  style: TextStyle(
                    fontFamily: 'Roboto',
                    fontWeight: FontWeight.w500,
                    fontStyle: FontStyle.normal,
                    fontSize: 12,
                    color: offlineColor,
                  ),
                ),
                SizedBox(height: 5,),
                SelectableText(
                  'hmac',
                  style: TextStyle(
                    fontFamily: 'Roboto',
                    fontWeight: FontWeight.w600,
                    fontStyle: FontStyle.normal,
                    fontSize: 12,
                    color: regularFontColor,
                  ),
                ),
                SizedBox(height: 5,),
                SelectableText(
                  'David9823124',
                  style: TextStyle(
                    fontFamily: 'Roboto',
                    fontWeight: FontWeight.w600,
                    fontStyle: FontStyle.normal,
                    fontSize: 12,
                    color: regularFontColor,
                  ),
                ),
                SizedBox(height: 5,),
                SelectableText(
                  '**********',
                  style: TextStyle(
                    fontFamily: 'Roboto',
                    fontWeight: FontWeight.w600,
                    fontStyle: FontStyle.normal,
                    fontSize: 12,
                    color: regularFontColor,
                  ),
                ),
                SizedBox(height: 20,),
                SelectableText(
                  'Endpoint',
                  style: TextStyle(
                    fontFamily: 'Roboto',
                    fontWeight: FontWeight.w500,
                    fontStyle: FontStyle.normal,
                    fontSize: 12,
                    color: offlineColor,
                  ),
                ),
                SizedBox(height: 5,),
                SelectableText(
                  'https://qingstor.com:443',
                  style: TextStyle(
                    fontFamily: 'Roboto',
                    fontWeight: FontWeight.w600,
                    fontStyle: FontStyle.normal,
                    fontSize: 12,
                    color: regularFontColor,
                  ),
                ),
              ],
            )
          ],
        ),
      ),
      onClose: controller.closeDialog,
      buttons: [],
    );
  }
}
