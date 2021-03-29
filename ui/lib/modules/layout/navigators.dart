import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../routes/index.dart';
import '../../widgets/side_link/index.dart';

class Navigators extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Container(
      width: 220,
      height: double.infinity,
      color: Colors.white,
      padding: EdgeInsets.symmetric(vertical: 20.0),
      child: Column(
        children: [
          Container(
            alignment: Alignment.centerLeft,
            margin: EdgeInsets.only(bottom: 48.0),
            padding: EdgeInsets.symmetric(horizontal: 20.0),
            child: SelectableText(
              'Project name'.tr,
              style: Theme.of(context).textTheme.headline4,
            ),
          ),
          SideLink(title: "Tasks", icon: Icons.dashboard, path: Routes.main),
          // SideLink(title: "Endpoints", icon: Icons.album, path: Routes.endpoints),
          // SideLink(title: "Agents", icon: Icons.source, path: "/agents"),
        ],
      )
    );
  }
}