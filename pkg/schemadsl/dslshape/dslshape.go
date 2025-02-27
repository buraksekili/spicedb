//go:generate stringer -type=NodeType -output zz_generated.nodetype_string.go

// Package dslshape defines the types representing the structure of schema DSL.
package dslshape

// NodeType identifies the type of AST node.
type NodeType int

const (
	// Top-level
	NodeTypeError   NodeType = iota // error occurred; value is text of error
	NodeTypeFile                    // The file root node
	NodeTypeComment                 // A single or multiline comment

	NodeTypeDefinition // A definition.

	NodeTypeRelation   // A relation
	NodeTypePermission // A permission

	NodeTypeTypeReference         // A type reference
	NodeTypeSpecificTypeReference // A reference to a specific type.

	NodeTypeUnionExpression
	NodeTypeIntersectExpression
	NodeTypeExclusionExpression

	NodeTypeArrowExpression // A TTU in arrow form.

	NodeTypeIdentifier // An identifier under an expression.
)

const (
	//
	// All nodes
	//
	// The source of this node.
	NodePredicateSource = "input-source"

	// The rune position in the input string at which this node begins.
	NodePredicateStartRune = "start-rune"

	// The rune position in the input string at which this node ends.
	NodePredicateEndRune = "end-rune"

	// A direct child of this node. Implementations should handle the ordering
	// automatically for this predicate.
	NodePredicateChild = "child-node"

	//
	// NodeTypeError
	//

	// The message for the parsing error.
	NodePredicateErrorMessage = "error-message"

	//
	// NodeTypeComment
	//

	// The value of the comment, including its delimeter(s)
	NodeCommentPredicateValue = "comment-value"

	//
	// NodeTypeDefinition
	//

	// The name of the definition
	NodeDefinitionPredicateName = "definition-name"

	//
	// NodeTypeRelation + NodeTypePermission
	//

	// The name of the relation/permission
	NodePredicateName = "relation-name"

	//
	// NodeTypeRelation
	//

	// The allowed types for the relation.
	NodeRelationPredicateAllowedTypes = "allowed-types"

	//
	// NodeTypeTypeReference
	//

	// A type under a type reference.
	NodeTypeReferencePredicateType = "type-ref-type"

	//
	// NodeTypeSpecificTypeReference
	//

	// A type under a type reference.
	NodeSpecificReferencePredicateType = "type-name"

	// A relation under a type reference.
	NodeSpecificReferencePredicateRelation = "relation-name"

	//
	// NodeTypePermission
	//

	// The expression to compute the permission.
	NodePermissionPredicateComputeExpression = "compute-expression"

	//
	// NodeTypeIdentifer
	//

	// The value of the identifier.
	NodeIdentiferPredicateValue = "identifier-value"

	//
	// NodeTypeUnionExpression + NodeTypeIntersectExpression + NodeTypeExclusionExpression + NodeTypeArrowExpression
	//
	NodeExpressionPredicateLeftExpr  = "left-expr"
	NodeExpressionPredicateRightExpr = "right-expr"
)
