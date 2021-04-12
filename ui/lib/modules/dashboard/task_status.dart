import 'package:flutter/material.dart';

import '../../common/colors.dart';

class TaskStatus extends StatelessWidget {
  final String status;

  Color get bgColor {
    if (status == "created") {
      return rgba(255, 206, 52, 1);
    } else if (status == "stopped") {
      return rgba(202, 38, 33, 1);
    } else if (status == "finished") {
      return rgba(0, 170, 114, 1);
    }
  }

  IconData get iconName {
    if (status == "created") {
      return Icons.more_horiz;
    } else if (status == "stopped") {
      return Icons.pause;
    } else if (status == "finished") {
      return Icons.done;
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
            width: 12,
            height: 12,
            alignment: Alignment.center,
            decoration: BoxDecoration(
              color: bgColor,
              borderRadius: BorderRadius.circular(6.0),
            ),
            child: Icon(
              iconName,
              size: 8,
              color: Colors.white,
            ),
          ),
          Padding(
            padding: EdgeInsets.only(left: 5),
            child: SelectableText(status),
          ),
        ],
      ),
    );
  }
}
