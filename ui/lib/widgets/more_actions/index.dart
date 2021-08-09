import 'package:get/get.dart';
import 'package:flutter/material.dart';

import '../../common/colors.dart';

class MoreActions extends StatelessWidget {
  final PopupMenuItemSelected<String> onSelected;
  final PopupMenuItemBuilder<String> itemBuilder;
  final bool enabled;

  MoreActions({
    required this.onSelected,
    required this.itemBuilder,
    this.enabled = true,
  });

  @override
  Widget build(BuildContext context) {
    return PopupMenuButton(
      onSelected: onSelected,
      itemBuilder: itemBuilder,
      enabled: enabled,
      tooltip: "More actions".tr,
      offset: Offset(0, 36),
      child: IconButton(
        icon: Icon(
          Icons.more_vert,
          color: regularFontColor,
        ),
        iconSize: 18,
        onPressed: null,
      ),
    );
  }
}
