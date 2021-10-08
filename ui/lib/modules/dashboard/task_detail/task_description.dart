import 'package:get/get.dart';
import 'package:ui/common/colors.dart';
import 'package:ui/common/svg_provider.dart';
import 'package:ui/models/task.dart';
import 'package:flutter/material.dart';
import 'package:ui/widgets/button/constants.dart';
import 'package:ui/widgets/button/index.dart';
import 'package:ui/widgets/dot/index.dart';
import 'package:ui/widgets/page_container/index.dart';
import 'package:ui/widgets/progress_bar/index.dart';
import 'package:ui/widgets/svg_Dot/index.dart';

import '../controller.dart';

class TaskDescription extends GetView<DashboardController> {

  @override
  Widget build(BuildContext context) {
    return WidgetContainer(
      margin: EdgeInsets.only(bottom: 16.0),
      child: Padding(
        padding: EdgeInsets.all(20.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.start,
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Container(
                  width: 60.0,
                  height: 60.0,
                  margin: EdgeInsets.only(right: 16.0),
                  decoration: new BoxDecoration(
                    color: rgba(104, 131, 237, 0.1),
                    borderRadius: BorderRadius.only(
                      topLeft: Radius.circular(60.0),
                      bottomLeft: Radius.circular(60.0),
                      bottomRight: Radius.circular(60.0),
                    ),
                  ),
                  child: Center(
                    child: Icon(IconData(0xe600, fontFamily: 'tpIcon'), size: 32, color: primaryBackgroundColor),
                  ),
                ),
                Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    SelectableText('DM Task - 1', style: Theme.of(context).textTheme.headline2),
                    SelectableText('task ID：DM 2021041223', style: Theme.of(context).textTheme.bodyText2),
                  ],
                ),
                SizedBox(width: 30,),
                Expanded(
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.start,
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      Dot(
                        dotTitle: "Running".tr,
                        dotColor: primaryColor,
                      ),
                      SizedBox(width: 20,),
                      SvgDot(
                        dotTitle: 'Agent 1',
                        src: 'images/agents.svg',
                      ),
                      SizedBox(width: 20,),
                      SvgDot(
                        dotTitle: 'Agent 2',
                        src: 'images/agents.svg',
                      ),
                    ],
                  ),
                ),
                Button(
                  icon: IconData(0xe605, fontFamily: 'tpIcon'),
                  type: ButtonType.primary,
                  child: Text('Pause'),
                  onPressed: () {},
                ),
              ],
            ),
            /// task progress
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Container(
                  width: 60.0,
                  height: 60.0,
                  margin: EdgeInsets.only(right: 16.0),
                ),
                Expanded(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      SelectableText(
                        'Task Progress'.tr,
                        style: TextStyle(
                          fontFamily: 'Roboto',
                          fontStyle: FontStyle.normal,
                          fontWeight: FontWeight.w700,
                          fontSize: 12,
                          color: headlineFontColor,
                        ),
                      ),
                      SizedBox(height: 14,),
                      ProgressBar(
                        ratio: 0.4,
                        barWidth: 480,
                        barHeight: 12,
                        barColor: onlineColor,
                        description: 'File size : 312 MB / 912 MB',
                      )
                    ],
                  ),
                ),
                Expanded(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      SelectableText(
                        'Used time : 9h 33min'.tr,
                        style: TextStyle(
                          fontFamily: 'Roboto',
                          fontStyle: FontStyle.normal,
                          fontWeight: FontWeight.normal,
                          fontSize: 12,
                          color: headlineFontColor,
                        ),
                      ),
                      SizedBox(height: 5,),
                      SelectableText(
                        'Remaining time : 10h 26min'.tr,
                        style: TextStyle(
                          fontFamily: 'Roboto',
                          fontStyle: FontStyle.normal,
                          fontWeight: FontWeight.normal,
                          fontSize: 12,
                          color: headlineFontColor,
                        ),
                      ),
                      SizedBox(height: 5,),
                      SelectableText(
                        'Number of files : 3525 / 9931'.tr,
                        style: TextStyle(
                          fontFamily: 'Roboto',
                          fontStyle: FontStyle.normal,
                          fontWeight: FontWeight.normal,
                          fontSize: 12,
                          color: headlineFontColor,
                        ),
                      ),
                    ],
                  ),
                )
              ],
            )
          ],
        ),
      ),
    );
  }
}
