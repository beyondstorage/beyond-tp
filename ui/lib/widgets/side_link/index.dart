import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:get/get.dart';
import 'package:flutter/cupertino.dart';

import '../../routes/index.dart';
import '../../common/colors.dart';

class SideLinkController extends GetxController {
  final RxBool hovered = false.obs;
}

class SideLink extends StatelessWidget {
  final String title;
  final String path;
  final IconData icon;

  SideLink({
    this.title,
    this.icon,
    @required this.path,
  });

  bool get isCurrentPage {
    if (path == Routes.main) {
      return Get.routing.current == path;
    }

    return Get.routing.current.indexOf(path) == 0;
  }

  @override
  Widget build(BuildContext context) {
    final SideLinkController c = SideLinkController();

    return GestureDetector(
      onTap: () => Get.toNamed(path),
      child: MouseRegion(
        cursor: SystemMouseCursors.click,
        onExit: (PointerExitEvent e) => c.hovered(false),
        onEnter: (PointerEnterEvent e) => c.hovered(true),
        child: Obx(() => Container(
          height: 32.0,
          alignment: Alignment.centerLeft,
          padding: EdgeInsets.symmetric(horizontal: 20),
          decoration: BoxDecoration(
            border: Border(
              left: BorderSide(
                width: 2.0,
                color: c.hovered.value || isCurrentPage
                  ? Theme.of(context).primaryColor
                  : Colors.transparent,
              )
            ),
            color: c.hovered.value || isCurrentPage
              ? rgba(245, 247, 250, 0.5) : Colors.transparent
          ),
          child: Row(
            children: [
              Icon(
                icon,
                size: 16,
                color: c.hovered.value || isCurrentPage
                  ? Theme.of(context).primaryColor
                  : Theme.of(context).textTheme.headline6.color
              ),
              Padding(
                padding: EdgeInsets.only(left: 8.0),
                child: Text(
                  title.tr,
                  style: c.hovered.value || isCurrentPage
                    ? Theme.of(context).primaryTextTheme.headline6
                    : Theme.of(context).textTheme.headline6,
                ),
              )
            ],
          ),
        )),
      ),
    );
  }
}