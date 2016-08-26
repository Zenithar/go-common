package neo4j

const (
	// Delete queries
	nodeDeleteByID     = `MATCH (n) WHERE id(n) = { id } OPTIONAL MATCH (n)-[r]-() DELETE r, n`
	nodeDeleteAllByIDs = `MATCH (n) WHERE id(n) in { ids } OPTIONAL MATCH (n)-[r]-() DELETE r, n`
	nodePurgeDatabase  = `MATCH (n) OPTIONAL MATCH (n)-[r]-() DELETE r, n`
	nodeDeleteByType   = "MATCH (n:`%s`) OPTIONAL MATCH (n)-[r]-() DELETE r, n"
	relDeleteByID      = `MATCH (n)-[r]->() WHERE ID(r) = { id } DELETE r`
	relDeleteAllByIDs  = `MATCH (n)-[r]->() WHERE id(r) in { ids } DELETE r`
	relPurgeDatabase   = `MATCH (n) OPTIONAL MATCH (n)-[r]-() DELETE r, n`
	relDeleteByType    = "MATCH (n)-[r:`%s`]-() DELETE r"

	// Aggreate queries
	nodeCountByType = "MATCH (n%s) RETURN COUNT(n)"

	// zero depth Finder queries
	nodeZDFindOneByID         = `MATCH (n) WHERE id(n) = { id } RETURN n`
	nodeZDFindAllByIDs        = `MATCH (n) WHERE id(n) in { ids } RETURN n`
	nodeZDFindAllByTypeAndIDs = "MATCH (n:`%s`) WHERE id(n) in { ids } RETURN n"
	nodeZDFindAllByType       = "MATCH (n:`%s`) RETURN n"
	nodeIDFindOneByID         = `MATCH (n) WHERE id(n) = { id } WITH n MATCH p=(n)-[*0..]-(m) RETURN p`
	nodeIDFindAllByIDs        = `MATCH (n) WHERE id(n) in { ids } WITH n MATCH p=(n)-[*0..]-(m) RETURN p`
	nodeIDFindAllByTypeAndIDs = "MATCH (n:`%s`) WHERE id(n) in { ids } WITH n MATCH p=(n)-[*0..]-(m) RETURN p"
	nodeIDFindAllByType       = "MATCH (n:`%s`) WITH n MATCH p=(n)-[*0..]-(m) RETURN p"
)
