import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../../models/task.dart';
import '../../../common/colors.dart';
import '../controller.dart';

class TaskDetailFormField extends StatelessWidget {
  final String label;
  final String value;

  TaskDetailFormField(this.label, this.value);

  String get _label {
    switch (label) {
      case "work_dir":
        return "Work dir";
        break;
      case "bucket_name":
        return "Bucket name";
        break;
      default:
        return label.capitalize;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.symmetric(horizontal: 30),
      margin: EdgeInsets.only(bottom: 16),
      child: Row(
        children: [
          SizedBox(
            width: 142,
            child: SelectableText(
              _label.tr,
              style: TextStyle(
                fontSize: 12,
                height: 1.67,
                color: Theme.of(context).primaryColorLight,
                fontWeight: FontWeight.w500,
              ),
            ),
          ),
          SelectableText(
            value,
            style: TextStyle(
              fontSize: 12,
              color: Theme.of(context).primaryColorLight,
            ),
          ),
        ],
      ),
    );
  }
}

class TaskDetailForm extends GetView<DashboardController> {
  @override
  Widget build(BuildContext context) {
    return Obx(() {
      var srcSet = controller.taskDetail.value?.storages?.isNotEmpty
          ? controller.taskDetail.value?.storages[0]
          : Storage.fromMap({});

      var dstSet = controller.taskDetail.value?.storages?.isNotEmpty
          ? controller.taskDetail.value?.storages[1]
          : Storage.fromMap({});

      return Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          TaskDetailFormField(
            "Task name",
            controller.taskDetail.value?.name ?? "-",
          ),
          SizedBox(height: 22),
          Container(
            decoration: new BoxDecoration(
              border: Border(
                bottom: BorderSide(
                  style: BorderStyle.solid,
                  color: rgba(228, 235, 241, 1),
                ),
              ),
            ),
          ),
          Padding(
            padding: EdgeInsets.only(top: 16, bottom: 24),
            child: SelectableText(
              "Source set".tr,
              style: TextStyle(
                fontSize: 12,
                height: 1.67,
                color: Colors.black,
                fontWeight: FontWeight.w600,
              ),
            ),
          ),
          TaskDetailFormField("Source type", srcSet?.type ?? "-"),
          ...(srcSet.options ?? []).map(
              (option) => TaskDetailFormField(option.key, option.value ?? "-")),
          SizedBox(
            height: 8,
          ),
          Container(
            decoration: new BoxDecoration(
              border: Border(
                bottom: BorderSide(
                  style: BorderStyle.solid,
                  color: rgba(228, 235, 241, 1),
                ),
              ),
            ),
          ),
          Padding(
            padding: EdgeInsets.only(top: 16, bottom: 24),
            child: SelectableText(
              "Destination set".tr,
              style: TextStyle(
                fontSize: 12,
                height: 1.67,
                color: Colors.black,
                fontWeight: FontWeight.w600,
              ),
            ),
          ),
          TaskDetailFormField("Destination type", dstSet?.type ?? "-"),
          ...(dstSet.options ?? []).map(
              (option) => TaskDetailFormField(option.key, option.value ?? "-")),
        ],
      );
    });
  }
}
