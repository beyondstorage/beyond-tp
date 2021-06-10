import 'package:flutter/rendering.dart';
import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../common/global.dart';
import '../../common/colors.dart';
import '../../routes/index.dart';
import '../../widgets/side_link/index.dart';

class Navigators extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Container(
      width: navigationWidth,
      height: double.infinity,
      color: rgba(52, 61, 190, 1),
      child: Column(
        children: [
          Padding(
            padding: EdgeInsets.symmetric(vertical: 32.0),
            child: Center(
              child: GestureDetector(
                onTap: () => Get.toNamed("/"),
                child: MouseRegion(
                  cursor: SystemMouseCursors.click,
                  child: Text('DM', style: TextStyle(
                    height: 1.0,
                    fontFamily: "Roboto",
                    fontSize: 32.0,
                    fontWeight: FontWeight.bold,
                    color: Colors.white,
                  )),
                ),
              ),
            ),
          ),
          SideLink(title: "Tasks".tr, icon: Icons.source, path: Routes.main),
          SideLink(title: "Agents".tr, icon: Icons.dns, path: Routes.agents),
          SideLink(title: "Identities".tr, icon: Icons.how_to_reg, path: Routes.identities),
        ],
      )
    );
  }
}
