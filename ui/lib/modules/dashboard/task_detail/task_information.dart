import 'package:date_format/date_format.dart';
import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:ui/common/colors.dart';
import 'package:ui/models/task.dart';
import 'package:ui/widgets/page_container/index.dart';
import 'package:ui/widgets/tabs/index.dart';
import 'package:ui/widgets/tabs/model.dart';

import 'time_log_pane.dart';
import 'base_information_pane.dart';

import '../controller.dart';


class TaskInformation extends GetView<DashboardController> {
  
  List<TabPane> TabsList = [
    TabPane(
      tabTitle: 'Base Information'.tr,
      pane: BaseInformationPane()
    ),
    TabPane(
      tabTitle: 'Real - Time Log'.tr,
      pane: TimeLogPane(
        logTitle: 'Unit Test (1.16, macos-latest)',
        description: 'Successed on 18 Mar in 53s',
        logList: [
          TimeLog(logContent: 'Set up job', time: 3),
          TimeLog(logContent: 'Set up Go 1.x', time: 2),
          TimeLog(logContent: 'Build', time: 44),
          TimeLog(logContent: 'Test', time: 23)
        ],
      )
    ),
  ];

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: WidgetContainer(
        child: Container(
          padding: EdgeInsets.all(10),
          child: Tabs(
            titleList: TabsList,
          ),
        ),
      ),
    );
  }
}
