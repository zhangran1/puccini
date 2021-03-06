// This file was auto-generated from a YAML file

package v1_0

func init() {
	Profile["/tosca/common/1.0/js/lib/utils.js"] = `

var tosca = {};

tosca.isTosca = function(o, kind) {
	if (o.metadata === undefined)
		return false;
	o = o.metadata['puccini-tosca'];
	if (o === undefined)
		return false;
	if (o.version !== '1.0')
		return false;
	if (kind !== undefined)
		return kind === o.kind;
	return true;
};

tosca.isNodeTemplate = function(vertex, typeName) {
	if (tosca.isTosca(vertex, 'NodeTemplate')) {
		if (typeName !== undefined)
			return typeName in vertex.properties.types;
		return true;
	}
	return false;
};

tosca.getNestedValue = function(singular, plural, args) {
	args = Array.prototype.slice.call(args);
	var length = args.length;
	if (length < 2)
		throw 'must have at least 2 arguments';
	var nodeTemplate = tosca.getModelableEntity(args[0]);
	var a = 1;
	var arg = args[a];
	var value = nodeTemplate[plural];
	if (arg in nodeTemplate.capabilities) {
		value = nodeTemplate.capabilities[arg][plural];
		singular = puccini.sprintf('capability "%s" %s', arg, singular);
		arg = args[++a];
	} else for (var r = 0, l = nodeTemplate.requirements.length; r < l; r++) {
		var requirement = nodeTemplate.requirements[r];
		if ((requirement.name === arg) && requirement.relationship) {
			value = requirement.relationship[plural];
			singular = puccini.sprintf('relationship "%s" %s', arg, singular);
			arg = args[++a];
			break;
		}
	}
	if (arg in value)
		value = value[arg];
	else
		throw puccini.sprintf('%s "%s" not found in "%s"', singular, arg, nodeTemplate.name);
	value = clout.coerce(value);
	for (var i = a + 1; i < length; i++) {
		arg = args[i];
		if (arg in value)
			value = value[arg];
		else
			throw puccini.sprintf('nested %s "%s" not found in "%s"', singular, args.slice(a, i+1).join('.'), nodeTemplate.name);
	}
	return value;
};

tosca.getModelableEntity = function(entity) {
	var vertex;
	switch (entity) {
	case 'SELF':
		if ((typeof $site === 'undefined') || !$site)
			throw puccini.sprintf('"%s" cannot be used in this context', entity);
		vertex = $site;
		break;
	case 'SOURCE':
		if ((typeof $source === 'undefined') || !$source)
			throw puccini.sprintf('"%s" cannot be used in this context', entity);
		vertex = $source;
		break;
	case 'TARGET':
		if ((typeof $target === 'undefined') || !$target)
			throw puccini.sprintf('"%s" cannot be used in this context', entity);
		vertex = $target;
		break;
	case 'HOST':
		if ((typeof $site === 'undefined') || !$site)
			throw puccini.sprintf('"%s" cannot be used in this context', entity);
		vertex = tosca.getHost($site);
		break;
	default:
		for (var vertexId in clout.vertexes) {
			var vertex = clout.vertexes[vertexId];
			if (tosca.isNodeTemplate(vertex) && (vertex.properties.name === entity))
				return vertex.properties;
		}
		vertex = {};
	}
	if (tosca.isNodeTemplate(vertex))
		return vertex.properties;
	else
		throw puccini.sprintf('node template "%s" not found', entity);
};

tosca.getHost = function(vertex) {
	for (var e = 0, l = vertex.edgesOut.length; e < l; e++) {
		var edge = vertex.edgesOut[e];
		if (tosca.isTosca(edge, 'Relationship')) {
			for (var typeName in edge.properties.types) {
				var type = edge.properties.types[typeName];
				if (type.metadata.role === 'host')
					return edge.target;
			}
		}
	}
	if (tosca.isNodeTemplate(vertex))
		throw puccini.sprintf('"HOST" not found for node template "%s"', vertex.properties.name);
	else
		throw '"HOST" not found';
};

tosca.getComparable = function(v) {
	if ((v === undefined) || (v === null))
		return null;
	var c = v.$number;
	if (c !== undefined)
		return c;
	c = v.$string;
	if (c !== undefined)
		return c;
	return v;
};

tosca.compare = function(v1, v2) {
	var c = v1.$comparer;
	if (c === undefined)
		c = v2.$comparer;
	if (c !== undefined)
		return clout.call(c, 'compare', [v1, v2]);
	v1 = tosca.getComparable(v1);
	v2 = tosca.getComparable(v2);
	if (v1 == v2)
		return 0;
	else if (v1 < v2)
		return -1;
	else
		return 1;
};
`
}
