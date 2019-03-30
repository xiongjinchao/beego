<aside class="main-sidebar">

    <!-- sidebar: style can be found in sidebar.less -->
    <section class="sidebar">

        <!-- Sidebar user panel (optional) -->
        <div class="user-panel">
            <div class="pull-left image">
                <img src="/static/AdminLTE/dist/img/user2-160x160.jpg" class="img-circle" alt="User Image">
            </div>
            <div class="pull-left info">
                <p>管理员</p>
                <!-- Status -->
                <a href="#"><i class="fa fa-circle text-success"></i> 在线</a>
            </div>
        </div>

        <!-- search form (Optional) -->
        <form action="#" method="get" class="sidebar-form">
            <div class="input-group">
                <input type="text" name="title" class="form-control" placeholder="搜索知识">
                <span class="input-group-btn">
              <button type="submit" id="search-btn" class="btn btn-flat"><i class="fa fa-search"></i>
              </button>
            </span>
            </div>
        </form>
        <!-- /.search form -->

        <!-- Sidebar Menu -->
        <ul class="sidebar-menu" data-widget="tree" data-current-controller="/home/">
            <li><a href="/"><i class="fa fa-th-large"></i> <span>系统面板</span></a></li>
            <li class="header">内容</li>
            <!-- Optionally, you can add icons to the links -->
            <li class="treeview">
                <a href="#">
                    <i class="fa fa-paper-plane"></i>
                    <span>知识管理</span>
                    <span class="pull-right-container">
                        <i class="fa fa-angle-left pull-right"></i>
                    </span>
                </a>
                <ul class="treeview-menu">
                    <li><a href="/knowledge-category"><i class="fa fa-th-list text-blue"></i> <span>知识目录</span></a></li>
                    <li><a href="/knowledge"><i class="fa fa-graduation-cap text-blue"></i> <span>知识管理</span></a></li>
                </ul>
            </li>
            <li class="treeview">
                <a href="#">
                    <i class="fa fa-wechat"></i>
                    <span>FAQ管理</span>
                    <span class="pull-right-container">
                        <i class="fa fa-angle-left pull-right"></i>
                    </span>
                </a>
                <ul class="treeview-menu">
                    <li><a href="/faq-category"><i class="fa fa-th-list text-red"></i> <span>FAQ分类</span></a></li>
                    <li><a href="/faq"><i class="fa fa-question-circle text-red"></i> <span>FAQ管理</span></a></li>
                </ul>
            </li>
            <li><a href="/notice"><i class="fa fa-rss"></i> <span>公告管理</span></a></li>
            <li class="header">系统</li>
            <li class="treeview">
                <a href="#">
                    <i class="fa fa-gears"></i>
                    <span>系统设置</span>
                    <span class="pull-right-container">
                        <i class="fa fa-angle-left pull-right"></i>
                    </span>
                </a>
                <ul class="treeview-menu">
                    <li><a href="/organization"><i class="fa fa-sitemap text-purple"></i> <span>组织架构</span></a></li>
                    <li><a href="/user"><i class="fa fa-github-alt text-purple"></i> <span>用户管理</span></a></li>
                    <li><a href="/role"><i class="fa fa-anchor text-purple"></i> <span>角色管理</span></a></li>
                </ul>
            </li>
            <li><a class="logout-link" href="/logout"><i class="fa fa-power-off text-red"></i> <span>退出登录</span></a></li>
        </ul>
        <!-- /.sidebar-menu -->
    </section>
    <!-- /.sidebar -->
</aside>