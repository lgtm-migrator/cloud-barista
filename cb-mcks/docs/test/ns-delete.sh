#!/bin/bash
# ------------------------------------------------------------------------------
# usage
if [ "$#" -lt 1 ]; then 
	echo "./ns-delete.sh <namespace>"
	echo "./ns-delete.sh cb-mcks-ns "
	exit 0
fi

source ./conf.env

# ------------------------------------------------------------------------------
# const

# ------------------------------------------------------------------------------
# variables

v_NAMESPACE="$1"
if [ "${v_NAMESPACE}" == "" ]; then read -e -p "namespace ? : "  v_NAMESPACE;	fi
if [ "${v_NAMESPACE}" == "" ]; then echo "[ERROR] missing <namespace>"; exit -1; fi


# ------------------------------------------------------------------------------
# print info.
echo ""
echo "[INFO]"
echo "- (Name of namespace)        is '${v_NAMESPACE}'"


# ------------------------------------------------------------------------------
# delete
delete() {
	curl -sX DELETE ${c_URL_TUMBLEBUG}/ns/${v_NAMESPACE} -H "${c_AUTH}" -H "${c_CT}" -o /dev/null -w "NAMESPACE.delete():%{http_code}\n"
}

# ------------------------------------------------------------------------------
# show
show() {
	echo "NAMESPACE_LIST";  curl -sX GET ${c_URL_TUMBLEBUG}/ns          -H "${c_AUTH}" -H "${c_CT}" | jq
}

# ------------------------------------------------------------------------------
if [ "$1" != "-h" ]; then 
	echo ""
	echo "------------------------------------------------------------------------------"
	delete; show;
fi
