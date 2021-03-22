import 'package:get/get.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import 'shared_prefs.dart';

import '../routes/index.dart';

Future<QueryResult> queryGraphQL(QueryOptions options) async {
  String server = await getConfig("server");

  if (server == null) {
    Get.toNamed(Routes.login);

    return QueryResult(source: QueryResultSource.optimisticResult);
  }

  GraphQLClient client = GraphQLClient(
    cache: GraphQLCache(), link: HttpLink("$server/graphql"));

  QueryResult result = await client.query(options);

  if (result.hasException) {
    // todo exception
    // print(result.exception.toString());
  }

  return result;
}
