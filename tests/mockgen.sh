#!/usr/bin/env bash
mockgen -destination=${GOPATH}/src/github.com/MongoDBNavigator/go-backend/tests/mock/system_info_reader.go -package=mock github.com/MongoDBNavigator/go-backend/domain/system/repository SystemInfoReader

mockgen -destination=${GOPATH}/src/github.com/MongoDBNavigator/go-backend/tests/mock/collection_reader.go -package=mock github.com/MongoDBNavigator/go-backend/domain/database/repository CollectionReader
mockgen -destination=${GOPATH}/src/github.com/MongoDBNavigator/go-backend/tests/mock/collection_writer.go -package=mock github.com/MongoDBNavigator/go-backend/domain/database/repository CollectionWriter

mockgen -destination=${GOPATH}/src/github.com/MongoDBNavigator/go-backend/tests/mock/database_reader.go -package=mock github.com/MongoDBNavigator/go-backend/domain/database/repository DatabaseReader
mockgen -destination=${GOPATH}/src/github.com/MongoDBNavigator/go-backend/tests/mock/database_writer.go -package=mock github.com/MongoDBNavigator/go-backend/domain/database/repository DatabaseWriter

mockgen -destination=${GOPATH}/src/github.com/MongoDBNavigator/go-backend/tests/mock/document_reader.go -package=mock github.com/MongoDBNavigator/go-backend/domain/database/repository DocumentReader
mockgen -destination=${GOPATH}/src/github.com/MongoDBNavigator/go-backend/tests/mock/document_writer.go -package=mock github.com/MongoDBNavigator/go-backend/domain/database/repository DocumentWriter

mockgen -destination=${GOPATH}/src/github.com/MongoDBNavigator/go-backend/tests/mock/index_reader.go -package=mock github.com/MongoDBNavigator/go-backend/domain/database/repository IndexReader
mockgen -destination=${GOPATH}/src/github.com/MongoDBNavigator/go-backend/tests/mock/index_writer.go -package=mock github.com/MongoDBNavigator/go-backend/domain/database/repository IndexWriter
