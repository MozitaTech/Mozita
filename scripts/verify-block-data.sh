# CELESTIA_APP_RPC="36657"

# # # Fetch the block data
# BLOCK_DATA=$(curl -s http://127.0.0.1:$CELESTIA_APP_RPC/block?height=1)

EVMOSD_RPC="36657"  # Replace with your evmosd RPC port if different

# Fetch the block data
BLOCK_DATA=$(curl -s http://localhost:$EVMOSD_RPC/block?height=1)

# Print the block data
echo $BLOCK_DATA