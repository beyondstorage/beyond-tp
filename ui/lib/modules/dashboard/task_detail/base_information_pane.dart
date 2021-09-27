import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:ui/common/colors.dart';

import 'identity_dialog.dart';

class BaseInformationPane extends StatelessWidget {
  const BaseInformationPane({ Key? key }) : super(key: key);

  Widget getTitle(String title) {
    return Row(
      children: [
        Container(
          margin: EdgeInsets.only(left: 6, right: 10, top: 10, bottom: 6),
          height: 4,
          width: 4,
          color: regularFontColor,
        ),
        SelectableText(
          title,
          style: TextStyle(
            fontFamily: 'Roboto',
            fontWeight: FontWeight.w600,
            fontStyle: FontStyle.normal,
            fontSize: 14,
            color: regularFontColor,
          ),
        )
      ],
    );
  }

  Widget getPane(List<Widget> children, {bool showBorder = true}) {
    return Container(
      padding: EdgeInsets.only(bottom: 30, top: 10),
      decoration: BoxDecoration(
        color: Colors.white,
        border: Border(bottom: BorderSide(color: showBorder ? rgba(226, 232, 240, 1) : Colors.white, width: 1, style: BorderStyle.solid)),
      ),
      child: Row(
        children: [
          Expanded(
            child: Column(
              children: children,
            ),
          )
        ],
      ),
    );
  }

  Widget getContent(String title, {String ?textContent, Widget ?customWidet}) {
    if (textContent == null && customWidet == null) {
      throw ArgumentError('Please fill in customWidet or textContent!');
    }
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        SelectableText(
          title,
          style: TextStyle(
            fontFamily: 'Roboto',
            fontWeight: FontWeight.w500,
            fontStyle: FontStyle.normal,
            fontSize: 12,
            color: offlineColor,
          ),
        ),
        SizedBox(height: 5,),
        textContent != null ? SelectableText(
          textContent as String,
          style: TextStyle(
            fontFamily: 'Roboto',
            fontWeight: FontWeight.w600,
            fontStyle: FontStyle.normal,
            fontSize: 12,
            color: regularFontColor,
          ),
        )
        : customWidet as Widget,
      ],
    );
  }

  Color getForeGroundColor(Set<MaterialState> states) {
    if (states.contains(MaterialState.hovered)) {
      return primaryColor;
    }

    return regularFontColor;
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.only(bottom: 30),
      child: Column(
        children: [
          Expanded(
            child: ListView(
              children: [
                getPane(
                  [
                    getTitle('source library settings'),
                    SizedBox(height: 20,),
                    Row(
                      children: [
                        SizedBox(width: 20,),
                        Expanded(
                          flex: 1,
                          child: getContent(
                            'target  library',
                            textContent: 'local files - FS',
                          ),
                        ),
                        Expanded(
                          flex: 2,
                          child: getContent(
                            'Work directory',
                            textContent: '/ home /  jack998 / Downloads',
                          ),
                        )
                      ],
                    )
                  ]
                ),
                getPane(
                  [
                    getTitle('target library settings'),
                    SizedBox(height: 20,),
                    Row(
                      children: [
                        SizedBox(width: 20,),
                        Expanded(
                          flex: 1,
                          child: getContent(
                            'target  library',
                            textContent: 'local files - FS'
                          ),
                        ),
                        Expanded(
                          flex: 2,
                          child: getContent(
                            'identity',
                            customWidet: TextButton(
                              style: ButtonStyle(
                                foregroundColor: MaterialStateProperty.resolveWith(getForeGroundColor),
                                overlayColor: MaterialStateProperty.resolveWith((states) => Colors.white),
                              ),
                              onPressed: () {
                                Get.dialog(IdentityDialog());
                              },
                              child: Text(
                                'QingStor - identity 1',
                                style: TextStyle(
                                  decoration: TextDecoration.underline,
                                )
                              )
                            )
                          ),
                        )
                      ],
                    ),
                    SizedBox(height: 20,),
                    Row(
                      children: [
                        SizedBox(width: 20,),
                        Expanded(
                          flex: 1,
                          child: getContent(
                            'bucket name',
                            textContent: 'QingStor bucket name',
                          ),
                        ),
                        Expanded(
                          flex: 2,
                          child: getContent(
                            'Work directory',
                            textContent: 'Root directory / file001 / ewqessad',
                          ),
                        )
                      ],
                    )
                  ]
                ),
                getPane(
                  [
                    getTitle('other settings'),
                    SizedBox(height: 20,),
                    Row(
                      children: [
                        SizedBox(width: 20,),
                        Expanded(
                          flex: 1,
                          child: getContent(
                            'task type',
                            textContent: 'One - time task',
                          ),
                        ),
                        Expanded(
                          flex: 2,
                          child: getContent(
                            'Data validate method',
                            textContent: 'md5',
                          ),
                        )
                      ],
                    ),
                    SizedBox(height: 20,),
                    Row(
                      children: [
                        SizedBox(width: 20,),
                        Expanded(
                          flex: 1,
                          child: getContent(
                            'Skip existing files condition',
                            textContent: 'file update time',
                          ),
                        )
                      ],
                    )
                  ],
                  showBorder: false,
                ),
              ],
            ),
          )
        ],
      ),
    );
  }
}
