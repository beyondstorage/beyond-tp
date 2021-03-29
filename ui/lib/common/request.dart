// import 'package:get/get.dart';
import 'package:flutter/foundation.dart' as Foundation;
import 'package:graphql_flutter/graphql_flutter.dart';

// import 'shared_prefs.dart';

// import '../routes/index.dart';

Future<QueryResult> queryGraphQL(QueryOptions options) async {
  // Support login page later
  // String server = await getConfig("server");
  //
  // if (server == null) {
  //   Get.toNamed(Routes.login);
  //
  //   return QueryResult(source: QueryResultSource.optimisticResult);
  // }

  String server = Foundation.kDebugMode ? "http://0.0.0.0:7436" : "";

  GraphQLClient client = GraphQLClient(
    cache: GraphQLCache(), link: HttpLink("$server/graphql"));

  QueryResult result = await client.query(options);

  if (result.hasException) {
    // todo exception
  }

  return result;
}
