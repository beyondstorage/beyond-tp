import 'package:flutter/material.dart';

import '../../common/colors.dart';

class TaskStatus extends StatelessWidget {
  final String status;

  Color get bgColor {
    switch (status) {
      case "Created":
        return rgba(241, 175, 78, 1);
      case "Ready":
        return rgba(241, 175, 78, 1);
      case "Running":
        return primaryColor;
      case "Stopped":
        return disableFontColor;
      case "Finished":
        return rgba(94, 191, 134, 1);
      default:
        return rgba(207, 59, 55, 1);
    }
  }

  String convertStatus(String status) {
    switch (status) {
      case "Created":
        return "To Be Run";
      case "Ready":
        return "To Be Run";
      case "Running":
        return "Running";
      case "Stopped":
        return "Paused";
      case "Finished":
        return "Completed";
      default:
        return "Running Error";
    }
  }

  TaskStatus(this.status);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.symmetric(vertical: 12.0, horizontal: 16.0),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Container(
            width: 8,
            height: 8,
            margin: EdgeInsets.only(top: 4),
            alignment: Alignment.center,
            decoration: BoxDecoration(
              color: bgColor,
              borderRadius: BorderRadius.circular(4),
            ),
          ),
          SizedBox(width: 8),
          SelectableText(
            convertStatus(status),
            style: TextStyle(
              fontSize: 12,
              color: regularFontColor,
            ),
          ),
        ],
      ),
    );
  }
}
