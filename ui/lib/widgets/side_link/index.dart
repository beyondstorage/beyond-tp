import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:get/get.dart';
import 'package:flutter/cupertino.dart';

import '../../routes/index.dart';
import '../../common/colors.dart';


class SideLink extends StatelessWidget {
  final String path;
  final String title;
  final IconData icon;

  SideLink({
    required this.title,
    required this.icon,
    required this.path,
  });

  bool get isCurrentPage {
    if (path == Routes.main) {
      return Get.routing.current == path;
    }

    return Get.routing.current.indexOf(path) == 0;
  }

  Color? getBackgroundColor(Set<MaterialState> states) {
    if (states.contains(MaterialState.hovered) || isCurrentPage) {
      return rgba(104, 131, 237, 1);
    }

    return null;
  }

  Color getForegroundColor(Set<MaterialState> states) {
    if (states.contains(MaterialState.hovered) || isCurrentPage) {
      return Colors.white;
    }

    return rgba(255, 255, 255, 0.7);
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.symmetric(horizontal: 12.0, vertical: 8.0),
      child: OutlinedButton(
        onPressed: () => Get.toNamed(path),
        style: ButtonStyle(
          backgroundColor: MaterialStateProperty.resolveWith(getBackgroundColor),
          foregroundColor: MaterialStateProperty.resolveWith(getForegroundColor),
          shape: MaterialStateProperty.resolveWith((states) => RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(8.0)
          )),
          padding: MaterialStateProperty.resolveWith((states) => EdgeInsets.zero),
          textStyle: MaterialStateProperty.resolveWith((states) => TextStyle(
            height: 1.75,
            fontFamily: "Roboto",
            fontSize: 12.0,
            fontWeight: FontWeight.w500,
          )),
        ),
        child: SizedBox(
          width: 56.0,
          height: 56.0,
          child: Center(
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [Icon(icon, size: 20), Text(title.tr)],
            ),
          ),
        ),
      ),
    );
  }
}