<template>
  <div class="user-guide">
    <!-- 顶部导航区 -->
    <div class="guide-header">
      <div class="header-content">
        <h1 class="title">
          <el-icon><QuestionFilled /></el-icon>
          小娇使用方法
        </h1>
        <div class="header-actions">
          <el-input
            v-model="searchKeyword"
            placeholder="搜索功能..."
            :prefix-icon="Search"
            clearable
            class="search-input"
          />
          <el-select v-model="categoryFilter" placeholder="功能分类" clearable class="category-select">
            <el-option label="全部功能" value="" />
            <el-option label="Spoke - 直接操作" value="spoke" />
            <el-option label="Hub - 汇总分析" value="hub" />
          </el-select>
          <el-radio-group v-model="viewMode" size="small">
            <el-radio-button value="grid"><el-icon><Grid /></el-icon></el-radio-button>
            <el-radio-button value="list"><el-icon><List /></el-icon></el-radio-button>
          </el-radio-group>
        </div>
      </div>
    </div>

    <!-- 主体内容区 -->
    <div class="guide-content">
      <!-- 左侧分类导航 -->
      <aside class="category-sidebar">
        <el-menu :default-active="activeCategory" @select="handleCategorySelect">
          <el-menu-item index="all">
            <el-icon><Menu /></el-icon>
            <span>全部功能</span>
          </el-menu-item>
          <el-sub-menu index="spoke">
            <template #title>
              <el-icon><Operation /></el-icon>
              <span>Spoke - 直接操作</span>
            </template>
            <el-menu-item v-for="cat in spokeCategories" :key="cat.id" :index="cat.id">
              {{ cat.name }}
            </el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="hub">
            <template #title>
              <el-icon><DataAnalysis /></el-icon>
              <span>Hub - 汇总分析</span>
            </template>
            <el-menu-item v-for="cat in hubCategories" :key="cat.id" :index="cat.id">
              {{ cat.name }}
            </el-menu-item>
          </el-sub-menu>
        </el-menu>
      </aside>

      <!-- 功能卡片区域 -->
      <main class="features-main">
        <div v-if="filteredFeatures.length === 0" class="empty-state">
          <el-empty description="未找到相关功能" />
        </div>
        <div v-else :class="['features-container', viewMode]">
          <div
            v-for="feature in filteredFeatures"
            :key="feature.id"
            class="feature-card"
            @click="showFeatureDetail(feature)"
          >
            <div class="card-header">
              <div class="icon-wrapper" :class="feature.type">
                <el-icon :size="24"><component :is="feature.icon" /></el-icon>
              </div>
              <div class="card-title-area">
                <h3 class="card-title">{{ feature.name }}</h3>
                <el-tag :type="feature.type === 'spoke' ? 'primary' : 'success'" size="small">
                  {{ feature.type === 'spoke' ? 'Spoke' : 'Hub' }}
                </el-tag>
                <el-tag v-if="feature.isNew" type="warning" size="small">新功能</el-tag>
              </div>
            </div>
            <p class="card-description">{{ feature.description }}</p>
            <div class="card-footer">
              <el-button type="primary" link size="small">
                <el-icon><Right /></el-icon>
                查看详情
              </el-button>
            </div>
          </div>
        </div>
      </main>

      <!-- 右侧快捷操作 -->
      <aside class="quick-actions">
        <el-card class="action-card">
          <template #header>
            <div class="card-header-text">
              <el-icon><Star /></el-icon>
              <span>常用功能</span>
            </div>
          </template>
          <div class="quick-links">
            <el-button v-for="item in popularFeatures" :key="item.id" text @click="quickNavigate(item)">
              {{ item.name }}
            </el-button>
          </div>
        </el-card>
      </aside>
    </div>

    <!-- 功能详情对话框 -->
    <el-dialog
      v-model="detailDialogVisible"
      :title="currentFeature?.name"
      width="800px"
      class="feature-detail-dialog"
    >
      <div v-if="currentFeature" class="feature-detail">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="功能类型">
            <el-tag :type="currentFeature.type === 'spoke' ? 'primary' : 'success'">
              {{ currentFeature.type === 'spoke' ? 'Spoke - 直接操作' : 'Hub - 汇总分析' }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="功能分类">
            {{ currentFeature.category }}
          </el-descriptions-item>
          <el-descriptions-item label="功能描述" :span="2">
            {{ currentFeature.description }}
          </el-descriptions-item>
          <el-descriptions-item label="使用场景" :span="2">
            {{ currentFeature.scenario }}
          </el-descriptions-item>
        </el-descriptions>

        <el-divider content-position="left">操作指引</el-divider>
        <div class="operation-guide">
          <el-steps direction="vertical" :active="currentFeature.steps?.length || 0">
            <el-step v-for="(step, index) in currentFeature.steps" :key="index" :title="step" />
          </el-steps>
        </div>

        <el-divider content-position="left">注意事项</el-divider>
        <el-alert
          v-for="(tip, index) in currentFeature.tips"
          :key="index"
          :title="tip"
          type="info"
          :closable="false"
          style="margin-bottom: 10px"
        />
      </div>
      <template #footer>
        <el-button @click="detailDialogVisible = false">关闭</el-button>
        <el-button v-if="currentFeature?.route" type="primary" @click="navigateToFeature">
          前往功能页面
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import {
  Search, Grid, List, Menu, Operation, DataAnalysis, QuestionFilled,
  Right, Star, Box, Monitor, Tools, Document, ChatDotRound,
  DataLine, Setting, Calendar, Warning, Coin, Files
} from '@element-plus/icons-vue'

const router = useRouter()

// 状态管理
const searchKeyword = ref('')
const categoryFilter = ref('')
const viewMode = ref('grid')
const activeCategory = ref('all')
const detailDialogVisible = ref(false)
const currentFeature = ref<any>(null)

// 分类定义
const spokeCategories = [
  { id: 'delivery', name: '交付物管理' },
  { id: 'cmdb', name: 'CMDB管理' },
  { id: 'test-env', name: '测试环境支持' },
  { id: 'ops-exec', name: '运维执行' },
  { id: 'inspection', name: '智能巡检' },
  { id: 'incident-spoke', name: '事件支持' }
]

const hubCategories = [
  { id: 'incident-hub', name: '事件统计' },
  { id: 'knowledge', name: '知识库' },
  { id: 'cpd', name: '合作方大屏' },
  { id: 'ai-summary', name: '智能总结' }
]

// 功能数据
const features = ref([
  // Spoke - 交付物管理
  { id: 1, name: '交付物物料查询', type: 'spoke', category: '交付物管理', icon: Box, description: '查询交付物的物料信息和版本详情', scenario: '需要查看某个交付物的详细物料清单时使用', steps: ['进入交付物管理页面', '输入交付物名称或ID', '点击查询按钮', '查看物料详情'], tips: ['确保输入正确的交付物标识', '支持模糊查询'], isNew: false },
  { id: 2, name: '交付物提交', type: 'spoke', category: '交付物管理', icon: Box, description: '提交新的交付物到系统', scenario: '开发完成后需要提交交付物进行审核', steps: ['填写交付物基本信息', '上传相关文件', '选择目标环境', '提交审核'], tips: ['确保所有必填项已填写', '文件格式需符合规范'], isNew: false },
  { id: 3, name: '版本投产计划查询', type: 'spoke', category: '交付物管理', icon: Calendar, description: '查询版本的投产计划和时间安排', scenario: '需要了解版本投产时间和计划安排', steps: ['进入投产计划页面', '选择版本号', '查看计划详情'], tips: ['关注投产时间窗口', '提前做好准备工作'], isNew: false },
  { id: 4, name: '交付物投产检视', type: 'spoke', category: '交付物管理', icon: Monitor, description: '检视交付物投产状态和结果', scenario: '投产后需要检查交付物状态', steps: ['选择投产批次', '查看投产状态', '确认投产结果'], tips: ['及时处理投产异常', '记录投产问题'], isNew: false },
  
  // Spoke - CMDB管理
  { id: 5, name: 'IP查询', type: 'spoke', category: 'CMDB管理', icon: Coin, description: '查询IP地址的详细信息和归属', scenario: '需要查找某个IP的详细信息和用途', steps: ['进入CMDB管理', '选择IP查询', '输入IP地址', '查看详细信息'], tips: ['支持IP段查询', '可查看IP使用历史'], isNew: false },
  { id: 6, name: 'DB查询', type: 'spoke', category: 'CMDB管理', icon: Coin, description: '查询数据库实例信息', scenario: '需要查找数据库连接信息或配置', steps: ['选择DB查询功能', '输入数据库名称', '查看配置信息'], tips: ['注意数据库访问权限', '保护敏感信息'], isNew: false },
  { id: 7, name: '业务线查询', type: 'spoke', category: 'CMDB管理', icon: DataLine, description: '查询业务线的组织架构和负责人', scenario: '需要了解业务线的组织结构', steps: ['进入业务线管理', '选择或搜索业务线', '查看详细信息'], tips: ['及时更新业务线信息'], isNew: false },
  { id: 8, name: '系统&子系统查询', type: 'spoke', category: 'CMDB管理', icon: Setting, description: '查询系统和子系统的配置信息', scenario: '需要了解系统架构和配置', steps: ['选择系统查询', '输入系统名称', '查看系统详情'], tips: ['注意系统依赖关系'], isNew: false },
  
  // Spoke - 运维执行
  { id: 9, name: '系统资源监控', type: 'spoke', category: '运维执行', icon: Monitor, description: '监控指定环境/IP的CPU/内存/磁盘使用情况', scenario: '需要实时查看服务器资源使用情况', steps: ['选择目标环境或IP', '选择监控指标', '查看实时数据', '设置告警阈值'], tips: ['定期检查资源使用趋势', '及时处理资源告警'], isNew: false },
  { id: 10, name: '值班人提醒', type: 'spoke', category: '运维执行', icon: Warning, description: '设置和管理值班人员提醒', scenario: '需要提醒值班人员关注重要事项', steps: ['进入值班管理', '选择值班人员', '设置提醒内容', '发送提醒'], tips: ['确保值班人员信息准确', '重要事项及时提醒'], isNew: false },
  
  // Spoke - 智能巡检
  { id: 11, name: '自动巡检', type: 'spoke', category: '智能巡检', icon: Tools, description: '配置和执行自动化巡检任务', scenario: '需要定期检查系统健康状态', steps: ['配置巡检任务', '设置巡检周期', '启动自动巡检', '查看巡检报告'], tips: ['合理设置巡检频率', '关注巡检异常'], isNew: true },
  { id: 12, name: '告警播报', type: 'spoke', category: '智能巡检', icon: Warning, description: '系统告警的自动播报和通知', scenario: '需要及时获知系统告警信息', steps: ['配置告警规则', '设置播报渠道', '启用告警播报'], tips: ['避免告警风暴', '及时处理告警'], isNew: true },
  
  // Hub - 事件统计
  { id: 13, name: '请求单统计', type: 'hub', category: '事件统计', icon: DataAnalysis, description: '统计和分析各类请求单数据', scenario: '需要分析请求单处理情况和趋势', steps: ['选择统计维度', '设置时间范围', '查看统计图表', '导出统计报告'], tips: ['定期分析请求单趋势', '优化处理流程'], isNew: false },
  { id: 14, name: '问题单统计', type: 'hub', category: '事件统计', icon: DataAnalysis, description: '统计和分析问题单处理情况', scenario: '需要了解问题单的分布和处理效率', steps: ['选择统计类型', '筛选问题单', '查看分析结果'], tips: ['关注高频问题', '总结解决方案'], isNew: false },
  
  // Hub - 知识库
  { id: 15, name: '知识库管理', type: 'hub', category: '知识库', icon: Document, description: '创建、查询和管理知识库文档', scenario: '需要沉淀和分享运维知识', steps: ['创建知识文档', '分类整理', '设置访问权限', '分享给团队'], tips: ['及时更新知识库', '规范文档格式'], isNew: false },
  
  // Hub - 合作方大屏
  { id: 16, name: '合作方大屏', type: 'hub', category: '合作方大屏', icon: Monitor, description: '展示合作方的用户概况、环境信息和运维状态', scenario: '需要向合作方展示运维数据', steps: ['配置大屏内容', '选择展示数据', '启动大屏展示'], tips: ['定期更新展示数据', '确保数据准确性'], isNew: false },
  
  // Hub - 智能总结
  { id: 17, name: '智能总结', type: 'hub', category: '智能总结', icon: ChatDotRound, description: '基于知识库的智能总结和问答', scenario: '需要快速获取知识库相关信息', steps: ['输入问题或关键词', '查看智能总结结果', '查看相关知识文档'], tips: ['问题描述要清晰', '善用知识库标签'], isNew: true }
])

// 热门功能
const popularFeatures = computed(() => features.value.slice(0, 6))

// 过滤后的功能列表
const filteredFeatures = computed(() => {
  let result = features.value

  // 分类过滤
  if (categoryFilter.value) {
    result = result.filter(f => f.type === categoryFilter.value)
  }

  // 类别过滤
  if (activeCategory.value !== 'all') {
    if (spokeCategories.some(c => c.id === activeCategory.value)) {
      result = result.filter(f => f.category === spokeCategories.find(c => c.id === activeCategory.value)?.name)
    } else if (hubCategories.some(c => c.id === activeCategory.value)) {
      result = result.filter(f => f.category === hubCategories.find(c => c.id === activeCategory.value)?.name)
    } else if (activeCategory.value === 'spoke') {
      result = result.filter(f => f.type === 'spoke')
    } else if (activeCategory.value === 'hub') {
      result = result.filter(f => f.type === 'hub')
    }
  }

  // 搜索过滤
  if (searchKeyword.value) {
    const keyword = searchKeyword.value.toLowerCase()
    result = result.filter(f =>
      f.name.toLowerCase().includes(keyword) ||
      f.description.toLowerCase().includes(keyword) ||
      f.category.toLowerCase().includes(keyword)
    )
  }

  return result
})

// 方法
const handleCategorySelect = (index: string) => {
  activeCategory.value = index
}

const showFeatureDetail = (feature: any) => {
  currentFeature.value = feature
  detailDialogVisible.value = true
}

const navigateToFeature = () => {
  if (currentFeature.value?.route) {
    router.push(currentFeature.value.route)
    detailDialogVisible.value = false
  }
}

const quickNavigate = (feature: any) => {
  showFeatureDetail(feature)
}
</script>

<style scoped>
.user-guide {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.guide-header {
  background: var(--gradient-card);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  padding: var(--spacing-lg) var(--spacing-xl);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-card);
  border: 1px solid var(--border-color);
  z-index: 10;
  margin-bottom: var(--spacing-lg);
  transition: all var(--transition-base);

  .header-content {
    max-width: 1600px;
    margin: 0 auto;
  }

  .title {
    font-size: 24px;
    font-weight: 600;
    color: #303133;
    margin: 0 0 16px 0;
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .header-actions {
    display: flex;
    gap: 12px;
    align-items: center;

    .search-input {
      width: 300px;
    }

    .category-select {
      width: 180px;
    }
  }
}

.guide-content {
  flex: 1;
  display: flex;
  gap: var(--spacing-lg);
  max-width: 1600px;
  margin: 0 auto;
  width: 100%;
  overflow: hidden;
}

.category-sidebar {
  width: 220px;
  background: var(--gradient-card);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-card);
  border: 1px solid var(--border-color);
  overflow-y: auto;
  padding: var(--spacing-md);
  transition: all var(--transition-base);
}

.features-main {
  flex: 1;
  overflow-y: auto;
  padding-right: 8px;
}

.features-container {
  &.grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
    gap: 16px;
  }

  &.list {
    display: flex;
    flex-direction: column;
    gap: 12px;
  }
}

.feature-card {
  background: var(--gradient-card);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-radius: var(--radius-xl);
  padding: var(--spacing-lg);
  box-shadow: var(--shadow-card);
  border: 1px solid var(--border-color);
  cursor: pointer;
  transition: all var(--transition-base);
  position: relative;
  overflow: hidden;

  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 3px;
    background: var(--gradient-accent);
    opacity: 0;
    transition: opacity var(--transition-base);
  }

  &:hover::before {
    opacity: 1;
  }

  &:hover {
    transform: translateY(-4px);
    box-shadow: var(--shadow-card-hover);
    border-color: var(--border-color-active);
  }

  .card-header {
    display: flex;
    gap: 12px;
    margin-bottom: 12px;
  }

  .icon-wrapper {
    width: 48px;
    height: 48px;
    border-radius: var(--radius-lg);
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
    box-shadow: var(--shadow-glow);

    &.spoke {
      background: var(--gradient-accent);
      color: white;
    }

    &.hub {
      background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
      color: white;
    }
  }

  .card-title-area {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 6px;
  }

  .card-title {
    font-size: 1rem;
    font-weight: 600;
    color: var(--text-primary);
    margin: 0;
  }

  .card-description {
    color: var(--text-muted);
    font-size: 0.875rem;
    line-height: 1.6;
    margin: 0 0 var(--spacing-md) 0;
  }

  .card-footer {
    display: flex;
    justify-content: flex-end;
  }
}

.quick-actions {
  width: 260px;
  flex-shrink: 0;

  .action-card {
    position: sticky;
    top: var(--spacing-lg);
    background: var(--gradient-card);
    backdrop-filter: blur(20px);
    -webkit-backdrop-filter: blur(20px);
    border-radius: var(--radius-xl);
    box-shadow: var(--shadow-card);
    border: 1px solid var(--border-color);
    transition: all var(--transition-base);
  }

  .card-header-text {
    display: flex;
    align-items: center;
    gap: 8px;
    font-weight: 600;
  }

  .quick-links {
    display: flex;
    flex-direction: column;
    gap: 8px;

    .el-button {
      justify-content: flex-start;
      padding: 8px 12px;
    }
  }
}

.empty-state {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
  background: var(--gradient-card);
  backdrop-filter: blur(20px);
  border-radius: var(--radius-xl);
  border: 1px solid var(--border-color);
}

.feature-detail {
  .operation-guide {
    margin: 20px 0;
  }
}

.feature-detail-dialog {
  :deep(.el-dialog__body) {
    max-height: 600px;
    overflow-y: auto;
  }
}
</style>
